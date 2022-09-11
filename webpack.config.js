const path = require('path');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

// webpack.config.js
module.exports = function (_env, argv) {
    const isProduction = argv.mode === "production";
    const isDevelopment = !isProduction;

    return {
        entry: ['./services/web/js/start.jsx'],
        output: {
            path: path.resolve(__dirname, '.build/'),
            filename: './web/public/js/bundle.js'
        },
        module: {
            rules: [
                {
                    parser: {
                        amd: false
                    }
                },
                {
                    test: /\.jsx?$/,
                    exclude: /node_modules/,
                    use: {
                        loader: "babel-loader",
                        options: {
                            cacheDirectory: true,
                            cacheCompression: false,
                            envName: isProduction ? "production" : "development"
                        }
                    }
                },
                {
                    test: /\.css$/,
                    use: [
                        isProduction ? MiniCssExtractPlugin.loader : "style-loader",
                        "css-loader"
                    ]
                }
            ]
        },
        resolve: {
            fallback: { "querystring": require.resolve("querystring-es3") },
            extensions: [".js", ".jsx"]
        },
        plugins: [
            isProduction &&
            new MiniCssExtractPlugin({
                filename: "assets/css/[name].[contenthash:8].css",
                chunkFilename: "assets/css/[name].[contenthash:8].chunk.css"
            })
        ].filter(Boolean)
    };
};