var myChart = echarts.init(document.getElementById('main'));

// 从后端接口获取K线数据
function fetchKLineData() {
    $.ajax({
        url: 'http://127.0.0.1:9217/kline?code=000001', // 替换为你的后端接口URL
        type: 'GET',
        dataType: 'json',
        success: function(data) {
            // 假设后端返回的数据格式是 [{open, close, lowest, highest, ...}, ...]
            var option = {
                // 配置tooltip
                tooltip: {
                    trigger: 'item', // 这里使用'item'触发类型，适用于K线图
                    //title: "000001",
                    triggerOn: 'mousemove', // 可以选择'mousemove'（鼠标移动时触发）或'click'（点击时触发）
                    formatter: function (params) {

                        // 自定义tooltip显示的内容
                        // params是包含了当前数据信息的对象
                        // 这里只是一个简单的例子，你可以根据需要进行修改
                        var date = params.name; // 时间戳或日期
                        var open = params.value[1]; // 开盘价
                        var close = params.value[2]; // 收盘价
                        var low = params.value[3]; // 最低价
                        var high = params.value[4]; // 最高价
                        return '日期: ' + date + '<br/>' +
                            '开盘价: ' + open + '<br/>' +
                            '收盘价: ' + close + '<br/>' +
                            '最低价: ' + low + '<br/>' +
                            '最高价: ' + high;
                    }
                },
                toolbox: {
                    show : true,
                    feature : {
                        mark : {
                            show : true,
                            title : {
                                mark : '辅助线-开关',
                                markUndo : '辅助线-删除',
                                markClear : '辅助线-清空'
                            },
                            lineStyle : {
                                width : 1,
                                color : '#1e90ff',
                                type : 'dashed'
                            }
                        },
                        dataZoom : {show: true},
                        dataView : {show: true, readOnly: false},
                        magicType: {show: true, type: ['line', 'bar']},
                        restore : {show: true},
                        saveAsImage : {show: true}
                    }
                },
                // ... 其他配置项 ...
                xAxis: {
                    type: 'category',
                    data: data.map(function(item, index) {
                        // 这里需要根据你的数据格式来提取时间戳或其他类别数据
                        return item.date; // 假设后端返回了时间戳
                    })
                },
                yAxis: {
                    scale: true,
                    splitArea: {
                        show: true
                    }
                },
                dataZoom: [
                    {
                        type: 'slider',
                        show: false,
                        start: 50, // 缩放滑块初始位置
                        end: 100 // 缩放滑块结束位置
                    },
                    {
                        type: 'inside', // 允许在图表内部进行缩放
                        start: 50,
                        end: 100
                    }
                ],
                series: [{
                    name: 'K线',
                    type: 'candlestick',
                    data: data.map(function(item) {
                        // 假设你的数据包含open, close, lowest, highest字段
                        return [item.open, item.close, item.low, item.high,item.date];
                    })
                }],
                grid: {
                    left: '0px',  // 左侧不留空隙
                    right: '0px',
                },
            };

            myChart.setOption(option);

            // 添加点击事件监听器
            myChart.on('click', function (params) {
                // 检查点击的是否是candlestick类型的数据
                if (params.componentType === 'series' && params.seriesType === 'candlestick') {
                    // 获取当前点击的蜡烛图的数据索引
                    var dataIndex = params.dataIndex;
                    // 假设你的数据集存储在option.series[0].data中
                    var candleData = option.series[0].data[dataIndex];

                    console.log(candleData)
                    // 你可以根据candleData中的信息来构造你要发送到后端的请求
                    var date = candleData[4]; // 假设第一个值是时间戳
                    var price = candleData[1];

                    var $form = $('#trade');
                    $form.find('input[name="date"]').val(date);
                    $form.find('input[name="price"]').val(price);

                    // 调用后端接口获取更多信息
                    //callBackendAPI(timestamp);
                }
            });
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.error('Error fetching data:', textStatus, errorThrown);
        }
    });
}



// 调用后端API获取更多信息
function callBackendAPI(timestamp) {
    alert(timestamp)
    $.ajax({
        url: '/your-backend-endpoint-for-more-info?timestamp=' + timestamp, // 替换为你的后端接口URL，并带上时间戳参数
        type: 'GET',
        dataType: 'json',
        success: function(data) {
            // 在这里处理从后端返回的数据
            console.log(data);
            // 例如，你可以更新图表上的某个提示框或显示一个模态框来显示这些数据
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.error('Error fetching data:', textStatus, errorThrown);
        }
    });
}


// 调用函数以获取数据并显示图表
fetchKLineData();


$("#buy").click(function (){
    alert("buy")
})

$("#sell").click(function (){
    alert("sell")
})


// 监听键盘事件
document.addEventListener('keydown', function(event) {
    xLen = myChart.getOption().xAxis[0].data.length
    var dataZoomOption = myChart.getOption().dataZoom[0]; // 获取数据区域缩放组件的配置
    var step = 100/xLen; // 每次移动的步长，可以根据需要调整
    switch (event.key) {
        case 'ArrowLeft': // 左键
            if (dataZoomOption.end > step) {
                dataZoomOption.start -= step;
                dataZoomOption.end -= step;
            }
            break;
        case 'ArrowRight': // 右键
            if (dataZoomOption.end < 100 - step) {
                dataZoomOption.start += step;
                dataZoomOption.end += step;
            }
        // 你可以添加对其他键的处理，如'ArrowUp'和'ArrowDown'进行缩放等
        default:
            break;
    }
    // 更新图表
    myChart.setOption({
        dataZoom: [dataZoomOption]
    });
});