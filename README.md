# GenFilesSyncExternalStoreAutomator

## インストール
1. [Releaseページ](https://github.com/nonaka-shu-lvgs/GenFilesSyncExternalStoreAutomator/releases/tag/v0.0.1)から実行ファイルをDLしてPATHが通っているディレクトリに置いてください。
2. 管理したいリポジトリのトップディレクトリにgenfiles.config.ymlという名前でファイルを作成して、設定ファイルを記述してください。
   - genfiles.config.yml.sampleにサンプルがあります。
3. リポジトリの.git/hooksディレクトリ内にpost-checkoutというファイル名でスクリプトを配置してください。
   - post-checkout.sampleを参考にしてください
   - 先ほどDLした実行ファイルを実行する内容であれば大丈夫です。