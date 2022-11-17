package postgres

import (
	"context"
	"screenShot-service/entity"
)

func (psql Postgres) CheckUrlExist(ctx context.Context, url string) (bool, error) {
	var status int

	err := conn.QueryRow(context.Background(), "select status from public.screenshot where url like $1  ", url).Scan(&status)
	if err != nil {
		return false, err
	}
	if status != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (psql Postgres) GetScreenShotByUrl(ctx context.Context, url string) (entity.ScreenShot, error) {
	var entity entity.ScreenShot

	err := conn.QueryRow(context.Background(), "select id,url,filepath,status from public.screenshot where url like $1  ", url).Scan(&entity.Id, &entity.Url, &entity.FilePath, &entity.Status)
	if err != nil {
		return entity, err
	}

	return entity, nil

}

func (psql Postgres) GetScreenShotById(ctx context.Context, id uint64) (entity.ScreenShot, error) {
	var entity entity.ScreenShot

	err := conn.QueryRow(context.Background(), "select id,url,filepath,status from public.screenshot where id like $1  ", id).Scan(&entity.Id, &entity.Url, &entity.FilePath, &entity.Status)
	if err != nil {
		return entity, err
	}

	return entity, nil

}

func (psql Postgres) CreateScreenShot(ctx context.Context, url string) (entity.ScreenShot, error) {
	var entity entity.ScreenShot

	err := conn.QueryRow(context.Background(), "insert into public.screenshot (url,status) values ($1,1) RETURNING id,url,filepath,status; ", url).Scan(&entity.Id, &entity.Url, &entity.FilePath, &entity.Status)

	// gg, err := conn.Exec(context.Background(), "insert into public.screenshot (url,status) values ($1,1); ", url)
	if err != nil {
		return entity, err
	}

	return entity, nil

}

// exists (
// 	select
// 		id
// 	from
// 		csf_manager_node
// 	where
// 		parent_id = node_id
// 		and done = false
// 		and id != node_id
// 	limit 1)
