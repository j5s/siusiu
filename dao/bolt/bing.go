package bolt

import (
	"biu/models"
)

func AddSearchResult(keyword, url string) error {
	return db.Save(&models.SearchResult{
		Keyword: keyword,
		URL:     url,
	})
}

func GetSearchResult(keyword string) (results []models.SearchResult, err error) {
	err = db.Find("Keyword", keyword, &results)
	return results, err
}

func GetAll() (results []models.SearchResult, err error) {
	err = db.AllByIndex("Keyword", &results)
	return results, err
}
