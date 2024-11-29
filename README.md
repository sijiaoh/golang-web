## このプロジェクト

Go言語によるWeb開発の学習のためのもの。

## プロジェクト構造

プロジェクト構造は[Standard Go Project Layout](https://github.com/golang-standards/project-layout)を参考にした。

## 採用フレームワーク

- [Air](https://github.com/air-verse/air)
    - ホットリロードのため。
- [Gin](https://github.com/gin-gonic/gin)
    - デファクトスタンダードっぽいため。
- [ent](https://github.com/go-gorm/gen)
    - Gormは[全削除問題](https://gorm.io/ja_JP/docs/delete.html)やV2が未だにリリースされていない問題から推測できる開発方針を嫌って回避。
    - sqlcは貧弱すぎる。
