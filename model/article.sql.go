package model

const (
	stmtQueryAllArticles = `
		SELECT
			id,
			COALESCE(title, '') AS title,
			COALESCE(content, '') AS content,
			COALESCE(thumbnail, '') AS thumbnail
		FROM
			articles`

	stmtQueryArticleByID = `
		SELECT
			id,
			COALESCE(title, '') AS title,
			COALESCE(content, '') AS content,
			COALESCE(thumbnail, '') AS thumbnail
		FROM
			articles
		WHERE
			id = ?`

	stmtQueryArticlesByTitle = `
		SELECT
			id,
			COALESCE(title, '') AS title,
			COALESCE(content, '') AS content,
			COALESCE(thumbnail, '') AS thumbnail
		FROM
			articles
		WHERE
			MATCH (title) AGAINST (? IN BOOLEAN MODE)`

	stmtInsertArticle = `
		INSERT INTO articles (title, content, thumbnail)
		VALUES (?, ?, ?)`

	stmtUpdateArticle = `
		UPDATE
			articles
		SET
			title = ?,
			content = ?,
			thumbnail = ?
		WHERE
			id = ?`
)
