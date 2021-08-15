package mysql

import (
	"bluebell01/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList()	(community []*models.Community, err error)  {
	sqlStr := "select community_id , community_name from community"
	if err := db.Select(&community,sqlStr) ; err != nil{
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
			community_id, community_name, introduction, create_time
			from community 
			where community_id = ?
	`
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}