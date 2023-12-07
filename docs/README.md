# 学生支援システム
このリポジトリは、名古屋工学院専門学校の学生が卒業制作として、学生支援システムである「NKCスマートドキュメント」略して「Nスマ」を開発するために用意したものです。

[English translation is here.](README_en.md)<br>
(Translated by DeepL)

# 使い方

### 必要なツール群
* Git
* Docker
* docker-compose

### Nスマを起動するには？

1. 上記ツール群を導入してください。

2. 開発環境で起動するには2-1のコマンドを、本番を想定した環境で起動するには2-2のコマンドを各種ターミナル(Git Bash推奨)にて実行してください。

2-1. 
```terminal
git clone https://github.com/nkc-cta20-team3/N-suma.git
cd N-suma
docker compose --profile dev_front --profile dev_back up -d --build
```

2-2. 
```terminal
git clone https://github.com/nkc-cta20-team3/N-suma.git
cd N-suma
docker compose -f docker-compose-preview.yaml up -d
```

3. Nスマを停止するとき、2-1のコマンドで起動した場合は3-1のコマンドを、3-2の場合は2のコマンドを実行してください。

3-1. 
``` terminal
docker compose stop   
```

3-2. 
``` terminal
docker compose -f docker-compose-preview.yaml stop
```

4. 完全に環境を破壊するには次のコマンドを実行してください。<br>
2-2のコマンドで構築した環境の場合は、fオプションでdocker-compose-preview.yamlファイルを指定するのを忘れないでください。
```
docker compose down --rmi all --volumes --remove-orphans
```

