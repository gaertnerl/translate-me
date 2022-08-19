const path = require('path');


var config = {
    mode: "production",
    devtool: false,
    resolve: {
        extensions: [".ts", ".tsx", ".js"],
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                loader: "ts-loader"
            }
        ]
    }
};

var gameConfig = Object.assign({}, config, {
    entry: {
        main: "./src/script/game.ts",
    },
    output: {
        path: path.resolve(__dirname, './src/compiled_js'),
        filename: "game.js"
    }
});

module.exports = [
    gameConfig,
];
