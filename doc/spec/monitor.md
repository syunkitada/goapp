# Monitor


## Logging
* ログの利用目的
    * ログの検索、閲覧
    * システム動作のトレーシング、リクエストのトレーシング
        * 開発者や運用者は、どこで処理が失敗したか、どこの処理に時間がかかったかをトレーシングする
        * ユーザは、自分のリクエストが処理されてるのか、失敗してるのかをトレーシングする
    * ログ解析によるアラート
        * 特定のログが発生した場合に、アラートを発生させ、緊急対応、日中対応する
* ログはトレーシングできるように、リクエストごとにidを発行し、伝搬させる
* 各デーモンサービスは、特定ディレクトリにログを出力する
* ローカルからも参照できるようにする


## Metrics
* メトリクスの利用目的
    * メトリクスの検索、閲覧
    * メトリクス解析によるアラート
* 各デーモンサービスは、特定ディレクトリにメトリクスを出力する
* sarのようにローカルからも参照できるようにする


## dataproxy-reporter
* reporterは、ログファイル、メトリクスファイルからデータを拾いdataproxyにデータを送る


## dataproxy
* log, tracelog, metricsをproxyする
* プロトコルはlineprotocolを参考にする
    * https://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/
    * log,index=cluster1,source=api log="create resource" 1465839830100400200
    * tracelog,index=cluster1,source=api,traceid=uuid1 msg="CreateNetworkV4(name=hoge,cluster=aaa)" 1465839830100400200
    * tracelog,index=cluster1,source=api,traceid=uuid1 msg="CreateNetworkV4(name=hoge,cluster=aaa): (err=nil)" 1465839830100400300
* proxyはデータをバッファに保存し、alertを設定してhookする
* indexにより、proxy先を変えて、シャーディングする
* データを検索するときもindex指定により、そのシャーディング先のDBから取得する
* traceidにより、ログを検索できるようにする


## alert-manager
* APIで受け取ったアラートをルールに従って、メールなどを配信する
* アラートの抑制はここで行う
