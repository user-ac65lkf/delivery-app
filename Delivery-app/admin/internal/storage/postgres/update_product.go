package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Shemistan/uzum_admin/internal/models"
	"log"
)

func (r *repo) UpdateProduct(ctx context.Context, req *models.Product) error {
	builder := sq.Update("products").SetMap(map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"price":       req.Price,
		"count":       req.Count,
	}).
		Where(sq.Eq{"id": req.ID}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	resp, err := builder.ExecContext(ctx)
	log.Printf("---> %+v", resp)
	if err != nil {
		return err
	}

	return nil
}

//
//func (r *repo) UpdateTwoProduct(ctx context.Context, req *models.Product) error {
//	tx, err := r.db.Beginx()
//	if err != nil {
//		log.Println("failed to create tranzaction")
//		return err
//	}
//
//	builder := sq.Insert("products").
//		Columns("name", "description", "price", "count").
//		Values(req.Name, req.Description, req.Price, req.Count).
//		RunWith(tx).
//		PlaceholderFormat(sq.Dollar).
//		Suffix(`RETURNING
//				"id"
//				`)
//
//	var id int64
//
//	err = builder.QueryRowContext(ctx).Scan(&id)
//	if err != nil {
//		errRollback := tx.Rollback()
//		if errRollback != nil {
//			log.Println()
//		}
//		return err
//	}
//
//	builder2 := sq.Insert("products").
//		Columns("name", "description", "price", "count").
//		Values(req.Name, req.Description, req.Price, req.Count).
//		RunWith(tx).
//		PlaceholderFormat(sq.Dollar).
//		Suffix(`RETURNING
//				"id"
//				`)
//
//	var id2 int64
//
//	err = builder2.QueryRowContext(ctx).Scan(&id2)
//	if err != nil {
//		errRollback := tx.Rollback()
//		if errRollback != nil {
//			log.Println()
//		}
//		return err
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		return err
//	}
//	return nil
//}
