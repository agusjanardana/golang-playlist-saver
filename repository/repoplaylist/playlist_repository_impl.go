package repoplaylist

import (
	"context"
	"playlist-saver/app/config/mysql"
	"playlist-saver/model/record"
)

type PlayListRepoImpl struct {
	client mysql.Client
}

func NewPlaylistRepository(client mysql.Client) PlaylistRepository {
	return &PlayListRepoImpl{client}
}

func (pl *PlayListRepoImpl) CreatePlaylist(ctx context.Context, name string, id int) (record.Playlist, error) {
	//panic("implement me")
	playlistRecord := record.Playlist{}
	playlistRecord.Name = name
	playlistRecord.UserId = id
	err := pl.client.Conn().Debug().WithContext(ctx).Create(&playlistRecord).Error
	if err != nil{
		return playlistRecord,err
	}

	return playlistRecord,nil
}

func (pl *PlayListRepoImpl) GetPlaylistByUserId(ctx context.Context, id int) ([]record.Playlist,error) {
	//panic("implement me")

	//playlistRecordArr := make([]record.Playlist, 0)
	var playlistRecord []record.Playlist
	//playlistRecord.Name = name
	//playlistRecord.UserId = id
	err := pl.client.Conn().Debug().WithContext(ctx).Where("user_id = ?", id).Find(&playlistRecord).Error
	if err != nil{
		return nil, err
	}

	return playlistRecord, nil
}
