package news

import (
	"testing"

	"github.com/SanderV1992/golang_simple_blog/database"
)

func TestFindAll(t *testing.T) {
	DB := database.Connect()
	repo := &RepoMysql{DB}

	_, count, err := repo.findAll(1, 5)
	if err != nil {
		t.Errorf("DB find error")
	}
	if count <= 0 {
		t.Errorf("DB 0-zero news")
	}
}

func TestFindByAlias(t *testing.T) {
	DB := database.Connect()
	repo := &RepoMysql{DB}

	_, err := repo.findByAlias("ddd_4")
	if err != nil {
		t.Errorf("DB full not found")
	}
}

func TestCreate(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}
