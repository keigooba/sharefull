# Sharefull
## 概要
既存のサービスをベースにして作成した求人サービスSharefull

基本的なCRUD機能に加え、チャット、画像アップロード、googleログインなどの機能を実装したサイトをHerokuへデプロイ

## 機能
1. 認証機能
   - ログイン、ログアウト
   - テストユーザーでログイン
   - Googleアカウントによるログイン(gomniauth)
2. 求人管理機能
   - 求人の作成、編集、削除
   - 現在日以降の求人情報の表示
   - 求人への応募
   - 応募中・募集中の求人表示
   - 応募中ユーザーの表示
3.ユーザー管理機能
   - 新規登録、編集、削除
   - 画像アップロード
4. Websocketを用いたチャット機能
   - 募集中求人へのメッセージ送信・保存
   - 応募状況での受信・送信メッセージの表示
   - 送信方法３種による画像表示（アップロード・Gravatar・Google）

## 技術
1. Go1.15.6
2. Heroku
3. postgres(開発ではsqlite3、本番ではpostgresを採用）
4. Bootstrap4.5.0
5. jQuery3.5.1

## 参考文献・作成物
1. Udemy【Go入門】Golang(Go言語)基礎〜応用 + 各種ライブラリ+ webアプリ開発コース(CRUD処理)  
   Todoリスト  
   https://github.com/keigooba/todo_app
2. オライリーの書籍「Go言語によるWebアプリケーション開発」  
   チャットアプリケーション・ドメイン名を検索するコマンドラインツール  
   https://github.com/keigooba/goblueprints
