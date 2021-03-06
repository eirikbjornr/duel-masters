package fx

import (
	"duel-masters/game/match"
	"fmt"
)

// ReturnToHand returns the card to the players hand instead of the graveyard
func ReturnToHand(card *match.Card, ctx *match.Context) {

	// When destroyed
	if event, ok := ctx.Event.(*match.CreatureDestroyed); ok {

		if event.Card == card {

			card.Player.MoveCard(card.ID, match.BATTLEZONE, match.HAND)
			ctx.Match.Chat("Server", fmt.Sprintf("%s was destroyed by %s and returned to the hand", event.Card.Name, event.Source.Name))

			ctx.InterruptFlow()

		}

	}

}

// ReturnToMana returns the card to the players manazone instead of the graveyard
func ReturnToMana(card *match.Card, ctx *match.Context) {

	// When destroyed
	if event, ok := ctx.Event.(*match.CreatureDestroyed); ok {

		if event.Card == card {

			card.Player.MoveCard(card.ID, match.BATTLEZONE, match.MANAZONE)
			card.Tapped = false

			ctx.Match.Chat("Server", fmt.Sprintf("%s was destroyed by %s but returned to the manazone", event.Card.Name, event.Source.Name))

			ctx.InterruptFlow()

		}

	}

}
