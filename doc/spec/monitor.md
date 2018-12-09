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
    * logやtracelogを解析して、metricsに変換してproxyする場合もある
* プロトコルはlineprotocolを参考にする
    * https://docs.influxdata.com/influxdb/v1.7/write_protocols/line_protocol_tutorial/
    * metrics
        * hoge,index=cluste1,host=host1 mem=123 1465839830100400200
    * log
        * log,index=cluster1,source=api,host=host1 msg="create resource" 1465839830100400200
    * tracelog
        * traceidにより、ログを検索できるようにする
        * tracelog,index=cluster1,source=api,host=host1,traceid=uuid1 msg="CreateNetworkV4(name=hoge,cluster=aaa)" 1465839830100400200
        * tracelog,index=cluster1,source=api,host=host1,traceid=uuid1 msg="CreateNetworkV4(name=hoge,cluster=aaa): (err=nil)" 1465839830100400300
* proxyはデータをバッファに保存し、alertを設定してhookする
    * 閾値やキーワードベースですぐにhookできるalertはここでhookする
* indexにより、proxy先を変えて、シャーディングする
    * 冗長化のためシャーディング先は2台から3台のDBがあると好ましい
    * シャーディング先に設定されたDBすべてに書き込む
        * 書き込めないのノードがいてもスルーする
    * データを検索するときは、index指定により、そのシャーディング先のDBすべてから検索し、データをマージして結果を返す
        * どこかのノードがデータを持っていればデータの冗長性が保たれる
    * シャーディング先で冗長化が担保できてるなら、書き込みも読み込みも1つから行う


## alert-manager
* APIで受け取ったアラートをルールに従って、メールなどを配信する
* アラートの抑制はここで行う
