package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

type Entity struct {
	id           EnumEntity
	restrictions map[int]EnumEntity
}

func (e Entity) checkRestrictions(list *[]Entity) (error, bool) {
	if len(*list) == 0 {
		return errors.New("No hay entidades para verificar"), false
	}

	for _, entity := range *list {
		if entity.id != NULLY {
			for _, restriction := range e.restrictions {

				if entity.id == restriction {
					return nil, false
				}
			}
		}
	}

	return nil, true
}

// State
type EnumEntity string

const (
	NULLY  EnumEntity = "🮻"
	CARROT EnumEntity = "🥕"
	WOLF   EnumEntity = "🐺"
	SHEEP  EnumEntity = "🐑"
)

var EntityList = []Entity{
	{CARROT, map[int]EnumEntity{0: SHEEP}},
	{WOLF, map[int]EnumEntity{0: SHEEP}},
	{SHEEP, map[int]EnumEntity{0: WOLF, 1: CARROT}},
}

var NULLYENTITY = Entity{NULLY, map[int]EnumEntity{}}

var title = color.New(color.FgCyan, color.Bold)
var subtitle = color.New(color.FgGreen, color.Bold)
var textBlue = color.New(color.FgBlue, color.Bold)
var textRed = color.New(color.FgHiRed)
var textYellow = color.New(color.FgYellow)

func main() {
	animalsRiver(EntityList)
}

func animalsRiver(list []Entity) {
	var EntityAside []Entity
	i := 0

	for i < 6 {
		title.Println("Iteración", i)
		subtitle.Println("Ida")
		checkRemoveInList(&list, &EntityAside)
		textBlue.Println("Side:", list)
		textBlue.Println("Other side:", EntityAside)
		subtitle.Println("Vuelta")
		checkRemoveInList(&EntityAside, &list)
		textBlue.Println("Side:", list)
		textBlue.Println("Other side:", EntityAside)

		if len(list) == 0 {
			break
		}

		i++
	}
}

func checkRemoveInList(side *[]Entity, otherSide *[]Entity) {
	for i, entity := range *side {
		if entity.id != NULLY {
			err, result := checkRemove(i, side)
			if err == nil && result == true {
				*otherSide = append(*otherSide, entity)
				(*side)[i] = NULLYENTITY
				fmt.Println("Se mueve ", entity.id)
				break
			} else {
				fmt.Println("No se puede mover", entity.id)
			}
		}
	}

	for i, entity := range *side {
		if entity.id == NULLY {
			*side = append((*side)[:i], (*side)[i+1:]...)
		}
	}
}

func checkRemove(pos int, entities *[]Entity) (error, bool) {
	if len(*entities) == 0 {
		return errors.New("No hay entidades para verificar"), false
	}

	if len(*entities) == 1 {
		return nil, false
	}

	if pos < 0 || pos >= len(*entities) {
		return errors.New("Posición inválida"), false
	}

	entitiesCp := make([]Entity, len(*entities))
	copy(entitiesCp, *entities)

	entitiesCp[pos] = NULLYENTITY

	textRed.Println("Verificando", entitiesCp)

	for i, entity := range entitiesCp {
		if pos != i {
			textYellow.Println(entity, "Escaneando")
			err, result := entity.checkRestrictions(&entitiesCp)

			if err == nil && result == false {
				return err, false
			}
		}
	}

	return nil, true
}
