// 基于准备好的dom，初始化echarts实例
var myChart = echarts.init(document.getElementById('main'));

// 指定图表的配置项和数据
var option = {
    title: {
        text: '小说作者影响力',
        subtext: '数据来自网络'
    },
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'shadow'
        }
    },

    grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
    },
    xAxis: {
        type: 'value',
        boundaryGap: [0, 0.01]
    },
    yAxis: {
        type: 'category',
        data: ['巴西','印尼','美国','印度','中国','世界人口(万)','美国','印度','中国']
    },
    series: [
        {
            name: '2011年',
            type: 'bar',
            itemStyle: {
                normal: {
                    color: function(params) {
                        //首先定义一个数组
                        var colorList = [
                            '#e8e3c0','#ef0064','#64BD3D','#EE9201','#29AAE3',
                            '#B74AE5','#0AAF9F','#E89589','#1dacaf','#c30e00'
                        ];
                        return colorList[params.dataIndex]
                    },
                    //以下为是否显示
                    label: {
                        show: false
                    }
                }
            },
            data: [18203, 23489, 29034, 104970, 231744, 690230,29034, 104970, 131744]
        },
    ]
};
