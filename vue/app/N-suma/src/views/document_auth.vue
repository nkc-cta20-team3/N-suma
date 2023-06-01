<template>
    <div class="title has-text-centered">
        各種書類認可
    </div>

    <div class="columns">
        
        <div class="column has-text-right">
            <div>未認可リスト</div>
        </div>
        
        <div class="column"></div>

        <div class="column">
            <div>ソート</div>
            <select class="select">
                <option>クラス略称</option>
                <option>名前</option>
                <option>分類</option>
            </select>
        </div>

    </div>

    <div class="container">
        <table class="table">
            <thead>
                <tr>
                    <th>クラス略称</th>
                    <th>氏名</th>
                    <th>分類</th>
                    <th>    </th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <th>CT4A</th>
                    <th>田中 太郎</th>
                    <th>資格</th>
                    <th><button>詳細</button></th>
                </tr>
                <tr>
                    <th>CT4B</th>
                    <th>田中 花子</th>
                    <th>就活</th>
                    <th><button>詳細</button></th>
                </tr>
                <tr>
                    <th>CT3A</th>
                    <th>鈴木 一郎</th>
                    <th>通常</th>
                    <th><button>詳細</button></th>
                </tr>
                <tr>
                    <th>CT3B</th>
                    <th>鈴木 次郎</th>
                    <th>通常</th>
                    <th><button>詳細</button></th>
                </tr>
                <tr>
                    <th>CT2A</th>
                    <th>パブロ・ディエゴ・ホセ・フランシスコ・デ・パウラ・ホアン・ネポムセーノ･マリーア・デ・ロス・レメディオス・クリスピン・クリスピアーノ・デ・ラ・サンディシマ・トリニダード･ルイス・イ・ピカソ</th>
                    <th>通常</th>
                    <th><button>詳細</button></th>
                </tr>
                <tr>
                    <th>CT2A</th>
                    <th>パブロ・ディエゴ・ホセ・フランシスコ・デ・パウラ・ホアン・ネポムセーノ･マリーア・デ・ロス・レメディオス・クリスピン・クリスピアーノ・デ・ラ・サンディシマ・トリニダード･ルイス・イ・ピカソ</th>
                    <th>通常</th>
                    <th><router-link to="/document_accept">詳細(aタグにしてます)</router-link></th>
                </tr>
            </tbody>
        </table>
    </div>
</template>