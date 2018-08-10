# Tracing


## Tracing Logger
* リクエストごとにidを発行し、各Actionに伝搬させる
* 各Actionは以下のフォーマットでログに出力する
    * trace;id;rpcsrc;action;msg;status;start;timespan
    * 例
        * trace;721b563f-ef1f-4248-9b3c-d8d0cbc8258d;;validatePassword;;2;201808101231;10
        * trace;721b563f-ef1f-4248-9b3c-d8d0cbc8258d;hoge.com;getAuthUser;;0;201808101231;10
        * trace;721b563f-ef1f-4248-9b3c-d8d0cbc8258d;;tokenIssue;user=admin,project=admin;2;201808101230;30


## MonitorAgent
* MonitorAgentはログを解析し、timespanが長い場合などで条件をかけアラートをMonitorControllerに送る
* MonitorContollerは、trace情報から親をたどり、状態を解析し、アラートを発生させる
