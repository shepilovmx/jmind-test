package repositories

import (
	"jmind-test/src/config"
	"jmind-test/src/models"
)

type BlockTotalsRepository struct{}

var BlockTotalsRep *BlockTotalsRepository

func (_ *BlockTotalsRepository) FindTotalsByBlockNumber(number int) (*models.BlockTotals, error) {
	blockTotals := &models.BlockTotals{}

	err := config.ServerCtx.Db.Model(blockTotals).
		Where("block_number = ?", number).
		Select()

	return blockTotals, err
}

func (_ *BlockTotalsRepository) InsertNewBlockTotals(blockTotals *models.BlockTotals) (*models.BlockTotals, error) {
	_, err := config.ServerCtx.Db.Model(blockTotals).Insert()

	return blockTotals, err
}
