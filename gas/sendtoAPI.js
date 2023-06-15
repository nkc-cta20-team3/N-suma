function sendDataToLambda() {
  // mainマージ時更新
  // スプレッドシートを取得
  //現在1番アクティブなスプレットシートを取得するため
  var sheet = SpreadsheetApp.getActiveSpreadsheet().getActiveSheet();
  
  // 現在スプレッドシートで一番下の行を取得
  var lastRow = sheet.getLastRow();
  var lastRowRange = sheet.getRange(lastRow, 1, 1, sheet.getLastColumn());
  
  // データを2次元配列として取得
  var data = lastRowRange.getValues()[0];
  
  // データをAPI Gatewayに送信するためのオブジェクトに変換
  var payload = {
    "value1": data[1],
    "value2": data[2],
    "value3": data[3],
    "value4": data[4],
    "value5": data[5],
    "value6": data[6],
    "value7": data[7],
    "value8": data[8],
    "value9": data[9],
    "value10": data[10]
  };

  // Lambdaエンドポイントにデータを送信するHTTPリクエストを作成
  var options = {
    "method": "post",
    "contentType": "application/json",
    "headers": {
      "x-api-key": "2fZVZfSYrLaBCMHcuLqJo92koQHFWnLg9aS4Gvtk" // 認証トークンを指定
    },
    "payload": JSON.stringify(payload)
  };
  
  // Lambdaエンドポイントにデータを送信
  var response = UrlFetchApp.fetch("https://jzmdclzhhgo5ck5xyid2newqwq0mpukn.lambda-url.us-east-1.on.aws/ ", options);
  
  // レスポンスをログに出力
  Logger.log(response.getContentText());
}