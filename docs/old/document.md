# 卒研メモ

## 画面設計

### ホーム画面
 - 各機能を呼び出す総合的な画面
 - 左側にサイドバーがある
 - (デザインイメージ)
 - 各種機能
     - 学生機能
         - 就職活動報告書の入力画面
             - フォームの入力画面
         - 個人データ管理画面
             - 自分の提出物を確認する画面
         - 資料閲覧画面
             - 提出された活動報告書などのデータを閲覧する画面
     - 教員機能
        - 書類認可画面
            - 提出された書類を処理する
                - 認可・拒否(再提出)
     - マスター管理者画面
        - 登録されたユーザーに教員フラグを付ける
  
### ログイン画面
 - 本来なら、SSO認証をここでする
 - ログインしていないときに遷移する画面
### アカウント登録画面
 - SSOを実装するまでのアカウント登録のつなぎ

## 技術選定について
現状、以下の技術を使用候補としている
 - php
 - mysql
 - docker
 - aws
使用言語として、phpを第一候補にしているが、以下の要素を考慮して別の言語を使用する可能性もある
 - クライアント側のレンダリング速度
 - クライアント側のメモリ使用率
 - サーバの負荷
 - 処理の応答速度
これらの要素を検証するには、クライアント側のアプリケーションにて、DBから取得した1000件のデータを読み込む処理を実行する。
理由として、システムの要件として全体のアカウントを読み込みする処理があることが確定しており、この処理が遅いとUXがかなり低下すると考えたため

# AWS使用サービス候補

## 基本的なサービス

### CostExplorer
サービスの使用量、利用状況を確認する

### AWS Budgets
一定以上の予算を超過した場合のアラーム

### VPC
絶対必要

### EC2
作業用サーバーとしても使用予定

### RDS
データベース

### S3
大規模ストレージ扱い

### Lamdba
サーバーレスアプリ

### Route53
ホストゾーンの登録

### Certificate Manager
SSL/TLS証明書のプロビジョニング


## 積極的に使いたいサービス

### IAM
CDK用のIAMユーザーの作成

### CloudFormation
CDKから環境を自動構築するのに使用

### CloudFront
CDNの提供

### Cloud9
クラウド開発環境の構築

### CloudShell
ブラウザベースのシェル

### Amazon Elastic Container Service (ECS)
コンテナデプロイ

### Amazon Elastic Container Registry(ECR)
コンテナイメージの保存

### Elastic Kubernetes Service(EKS)
kubernets

### AWS Cloud Map
クラウドの動的マップ

### API Gateway
APIの構築と管理

### Amazon EventBridge
イベント駆動型アーキテクチャを構築する

### Global Accelerator
アプリケーションへのトラフィックの最適化


## 追加機能用のサービス

### Amazon Simple Email Service
Eメールの送信

### Amazon Textract
ドキュメントからデータの抽出

### Amazon Polly
テキストの音声化

### Amazon Transcribe
音声の文字起こし

### Amazon Translate
言語の翻訳

## できれば使いたいサービス

### AWS Amplify
### AWS AppSync
### Amazon Managed BlockChain