package budget_realization

import (
	_ "github.com/lib/pq"
)

type BudgetRealization struct {
	Id                int    `form: "id" json: "id"`
	BudgetId          string `form: "budget_id" json: "budget_id"`
	ParentId          int32  `form: "parent_id" json: "parent_id"`
	Narration         string `form: "narration" json: "nareration"`
	Ref               string `form: "ref" json: "ref"`
	Date              string `form: "date" json: "date"`
	BudgetRealization string `form: "budget_realisasi" json: "budget_realisasi"`
}

type BudgetRealizationResponse struct {
	Status  int    `json: "status"`
	Message string `json: "message"`
	Data    []BudgetRealization
}
