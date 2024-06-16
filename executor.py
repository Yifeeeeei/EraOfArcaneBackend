from enum import Enum

import copy


class HandlerType(Enum):
    ENTER = 1
    DAMAGE = 3
    DIE = 5
    ADD_LIFE = 7


class TriggerType(Enum):
    BEFORE_ENTER = 1
    AFTER_ENTER = 2
    BEFORE_DAMAGE = 3
    AFTER_DAMAGE = 4
    BEFORE_DIE = 5
    AFTER_DIE = 6
    BEFORE_ADD_LIFE = 7
    AFTER_ADD_LIFE = 8


class Trigger:
    def __init__(self, triggerType: TriggerType, triggerFunction):
        self.triggerFunction = triggerFunction
        self.triggerType = triggerType

    def trigger(self, handlerParams, context, executor):
        return self.triggerFunction(handlerParams, context, executor)


class Handler:
    def __init__(self, handlerType: HandlerType, params):
        self.handlerType = handlerType
        self.params = params


class Executor:
    def __init__(self):
        self.maxIter = 100
        self.triggers = {}
        self.handlerQueue = []
        self.handlerPin = 0

        self.corresbondingTriggersBefore = {
            HandlerType.ENTER: TriggerType.BEFORE_ENTER,
            HandlerType.DAMAGE: TriggerType.BEFORE_DAMAGE,
            HandlerType.DIE: TriggerType.BEFORE_DIE,
            HandlerType.ADD_LIFE: TriggerType.BEFORE_ADD_LIFE,
        }
        self.correspondingTriggersAfter = {
            HandlerType.ENTER: TriggerType.AFTER_ENTER,
            HandlerType.DAMAGE: TriggerType.AFTER_DAMAGE,
            HandlerType.DIE: TriggerType.AFTER_DIE,
            HandlerType.ADD_LIFE: TriggerType.AFTER_ADD_LIFE,
        }

    def registerTrigger(self, trigger: Trigger) -> str:
        if trigger.triggerType not in self.triggers:
            self.triggers[trigger.triggerType] = []
        self.triggers[trigger.triggerType].append(trigger)

    def commitHandler(self, handler, position=None) -> ValueError:
        if position is None:
            self.handlerQueue.append(handler)
        elif position == "next":
            self.handlerQueue.insert(self.handlerPin + 1, handler)
        else:

            return "Invalid position"
        return None

    def execute(self, context: dict) -> tuple[dict, str]:
        while self.handlerPin < self.maxIter and self.handlerPin < len(
            self.handlerQueue
        ):

            handler = self.handlerQueue[self.handlerPin]
            handlerParams = handler.params

            triggerBefore = self.corresbondingTriggersBefore[handler.handlerType]
            if triggerBefore in self.triggers:
                for trigger in self.triggers[triggerBefore]:
                    handlerParams, context, error = trigger.trigger(
                        handlerParams, context, self
                    )
                    if error:
                        print(f"Error in iteration {self.handlerPin}: {error}")
                        return context, error

            if handler.handlerType == HandlerType.ENTER:
                print("Enter")
                context["cards"].append(handlerParams["card"])
            elif handler.handlerType == HandlerType.DIE:
                print("Die")
                for i, card in enumerate(context["cards"]):
                    if card["UUID"] == handlerParams["UUID"]:
                        context["cards"].pop(i)
                        break
            elif handler.handlerType == HandlerType.DAMAGE:
                print("Damage")
                value = handlerParams["value"]
                targetUUID = handlerParams["targetUUID"]
                for card in context["cards"]:
                    if card["UUID"] == targetUUID:
                        card["life"] -= value
                        if card["life"] <= 0:
                            self.commitHandler(
                                Handler(HandlerType.DIE, {"UUID": targetUUID}), "next"
                            )
                        break
            elif handler.handlerType == HandlerType.ADD_LIFE:
                print("Add Life")
                value = handlerParams["value"]
                targetUUID = handlerParams["targetUUID"]
                for card in context["cards"]:
                    if card["UUID"] == targetUUID:
                        card["life"] += value
                        break
            else:
                print("Unknown")

            triggerAfter = self.correspondingTriggersAfter[handler.handlerType]
            if triggerAfter in self.triggers:
                for trigger in self.triggers[triggerAfter]:
                    handlerParams, context, error = trigger.trigger(
                        handlerParams, context, self
                    )
                    if error:
                        print(f"Error in iteration {self.handlerPin}: {error}")
                        return context, error

            self.handlerPin += 1
        self.handlerQueue = []
        self.handlerPin = 0
        return context, None


# TEST
def test():
    context = {"cards": []}
    executor = Executor()

    # A has life = 2, B has (all damages + 1), C has 1 life and effect (when a unit dies, it gains 1 life), D has onEnter, deal 1 damage
    A_card = {"UUID": "A", "life": 2}
    # A enter
    executor.commitHandler(Handler(HandlerType.ENTER, {"card": A_card}))

    B_card = {"UUID": "B", "life": 1}
    # B enter
    executor.commitHandler(Handler(HandlerType.ENTER, {"card": B_card}))

    # register B's effect
    def B_effect(handlerParams, context, executor):
        newHanderParams = copy.deepcopy(handlerParams)
        newHanderParams["value"] += 1
        return newHanderParams, context, None

    executor.registerTrigger(Trigger(TriggerType.BEFORE_DAMAGE, B_effect))

    C_card = {"UUID": "C", "life": 1}
    # C enter
    executor.commitHandler(Handler(HandlerType.ENTER, {"card": C_card}))

    # register C's effect
    def C_effect(handlerParams, context, executor):
        executor.commitHandler(
            Handler(HandlerType.ADD_LIFE, {"targetUUID": "C", "value": 1}), "next"
        )
        return handlerParams, context, None

    executor.registerTrigger(Trigger(TriggerType.AFTER_DIE, C_effect))

    D_card = {"UUID": "D", "life": 1}
    context, _ = executor.execute(context)

    executor.commitHandler(Handler(HandlerType.ENTER, {"card": D_card}))
    context, _ = executor.execute(context)
    executor.commitHandler(
        Handler(HandlerType.DAMAGE, {"targetUUID": "A", "value": 1}), "next"
    )
    context, _ = executor.execute(context)
    print(context["cards"])


test()
