const CopyWebpackPlugin = require('copy-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    // This option controls if and how source maps are generated.
    // https://webpack.js.org/configuration/devtool/
    devtool: 'source-map',

    // https://webpack.js.org/concepts/entry-points/#multi-page-application
    entry: {
        index: './src/index.js'
    },

    // https://webpack.js.org/concepts/loaders/
    module: {
        rules: [
            // {
            //     test: /\.(eot|woff2?|ttf|svg)$/,
            //     use: [
            //         'file-loader'
            //     ]
            // },
            {
                test: /\.css$/,
                use: [
                    'style-loader',
                    "css-loader"
                ]
            },
            {
                test: /\.(eot|woff2?|ttf|svg)$/,
                use: [
                    {
                        loader: "url-loader", options: {
                            name: "[name]-[hash:5].min.[ext]", limit: 5000, // fonts file size <= 5KB, use 'base64'; else, output svg file
                        }
                    }
                ]
            },
            {
                test: /\.less$/,
                use: [
                    'style-loader',
                    'css-loader',
                    'less-loader'
                ]
            }
        ]
    },

    // https://webpack.js.org/concepts/plugins/
    plugins: [
        new CopyWebpackPlugin({
            patterns:[
                { from: 'src/images', to: 'images' },
            ],
        }),
        new HtmlWebpackPlugin({
            template: './src/index.html',
            inject: true,
            chunks: ['index'],
            title: 'Custom template',
            filename: 'index.html'
        })
    ],
};