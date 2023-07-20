package models

type Article struct {
	Id       uint16 `json:"id"`
	Title    string `json:"title"`
	Anons    string `json:"anons"`
	FullText string `json:"fullText"`
}

type UsersArticle struct {
	Id        int
	UserId    int
	ArticleId int
}
