(function(e){function t(t){for(var a,r,o=t[0],s=t[1],c=t[2],d=0,p=[];d<o.length;d++)r=o[d],Object.prototype.hasOwnProperty.call(i,r)&&i[r]&&p.push(i[r][0]),i[r]=0;for(a in s)Object.prototype.hasOwnProperty.call(s,a)&&(e[a]=s[a]);u&&u(t);while(p.length)p.shift()();return l.push.apply(l,c||[]),n()}function n(){for(var e,t=0;t<l.length;t++){for(var n=l[t],a=!0,o=1;o<n.length;o++){var s=n[o];0!==i[s]&&(a=!1)}a&&(l.splice(t--,1),e=r(r.s=n[0]))}return e}var a={},i={app:0},l=[];function r(t){if(a[t])return a[t].exports;var n=a[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,r),n.l=!0,n.exports}r.m=e,r.c=a,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)r.d(n,a,function(t){return e[t]}.bind(null,a));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],s=o.push.bind(o);o.push=t,o=o.slice();for(var c=0;c<o.length;c++)t(o[c]);var u=s;l.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},2649:function(e,t,n){"use strict";n("4740")},"356b":function(e,t,n){"use strict";n("e7c6")},4740:function(e,t,n){},"56d7":function(e,t,n){"use strict";n.r(t);n("cadf"),n("551c"),n("f751"),n("097d");var a=n("2b0e"),i=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("head-top"),n("div",[n("router-view")],1)],1)},l=[],r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"header_container"},[n("el-menu",{staticClass:"el-menu-demo",attrs:{"default-active":e.activeIndex,mode:"horizontal","background-color":"#545c64","text-color":"#fff","active-text-color":"#ffd04b",router:""},on:{select:e.handleSelect}},[n("el-menu-item",{attrs:{index:"/"}},[e._v("首页")]),n("el-menu-item",{attrs:{index:"/warn"}},[e._v("告警配置")]),n("el-menu-item",{attrs:{index:"/send"}},[e._v("发送配置")]),n("el-menu-item",{attrs:{index:"/warning"}},[e._v("告警信息")])],1)],1)},o=[],s={name:"HeadTop",data:function(){return{activeIndex:""}},created:function(){},computed:{},methods:{handleSelect:function(e,t){console.log(e,t)}}},c=s,u=(n("a9f6"),n("2877")),d=Object(u["a"])(c,r,o,!1,null,null,null),p=d.exports,h=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"hello"},[e._v("\n  hello\n")])},m=[],f={name:"HelloWorld",props:{msg:String}},v=f,b=Object(u["a"])(v,h,m,!1,null,"6a4dca19",null),g=b.exports,y={name:"app",components:{HelloWorld:g,HeadTop:p}},_=y,w=(n("7c55"),Object(u["a"])(_,i,l,!1,null,null,null)),k=w.exports,C=n("8c4f"),$=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home"},[e._m(0),n("div",{staticClass:"table-container"},[n("el-row",[n("el-col",{attrs:{span:18}},[n("div",[e._v(" ")])]),n("el-col",{attrs:{span:6}},[n("div",[n("el-input",{attrs:{size:"mini",placeholder:"输入区域或ip关键字搜索"},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})],1)])],1),n("el-table",{attrs:{data:e.tableData.filter((function(t){return!e.search}))}},[n("el-table-column",{attrs:{label:"Host",align:"center","header-aligh":"center",formatter:e.formateAgent}}),n("el-table-column",{attrs:{label:"区域",align:"center","header-aligh":"center",prop:"local"}}),n("el-table-column",{attrs:{formatter:e.formateLive,prop:"is_live",width:"100",align:"center","header-aligh":"center",label:"是否存活"}}),n("el-table-column",{attrs:{align:"center","header-aligh":"center","min-width":"200px",label:"agent管理"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(n){return e.handleWarn(t.row)}}},[e._v("告警规则管理")]),n("el-button",{attrs:{type:"warning",size:"mini"},on:{click:function(n){return e.handleSend(t.row)}}},[e._v("发送配置管理")]),n("el-button",{attrs:{type:"danger",size:"mini"},on:{click:function(n){return e.handleHistory(t.row)}}},[e._v("告警历史详情")]),n("el-button",{attrs:{type:"info",size:"mini"},on:{click:function(n){return e.handleShow(t.row)}}},[e._v("数据详情")])]}}])})],1)],1)])},x=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"headline"},[n("h3",[e._v(" 主机列表 ")])])}],L=n("ade3"),T=n("bc3a"),D=n.n(T),S="http://120.0.0.1:8083",E=n("4328"),j=n.n(E);function O(e,t){return new Promise((function(n,a){D.a.post(e,t).then((function(e){n(e)}),(function(e){a(e)})).catch((function(e){a(e)}))}))}function P(e,t){return new Promise((function(n,a){D.a.get(e,{params:t}).then((function(e){n(e)}),(function(e){a(e)})).catch((function(e){a(e)}))}))}D.a.defaults.timeout=5e3,D.a.defaults.headers.post["Content-Type"]="application/x-www-form-urlencoded;charset=UTF-8",D.a.defaults.baseURL=S,D.a.interceptors.request.use((function(e){return"post"===e.method&&(e.data=j.a.stringify(e.data)),e}),(function(e){return Promise.reject(e)})),D.a.interceptors.response.use((function(e){return 200==e.status&&"000000"==e.data.code?Promise.resolve(e.data):Promise.reject(e.data)}),(function(e){return Promise.reject({msg:e.message})}));var q;function z(){return P("/agent/list",{})}function I(){return P("/agent/sendList",{})}var V=(q={name:"home",data:function(){return{dialogVisible:!1,search:"",tableData:[{id:1,ip:"127.0.0.1",local:"beijing"},{id:2,ip:"127.0.0.2",local:"beijing"},{id:3,ip:"127.0.0.3",local:"beijing"}]}},methods:{deleteData:function(e,t){this.dialogVisible=!1,console.log(e),console.log(t),this.tableData.splice(e,1)},openDelete:function(e,t){var n=this;this.$confirm("此操作将永久删除该文件, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){n.deleteData(e,t),n.$message({type:"success",message:"删除成功!"})})).catch((function(){n.$message({type:"info",message:"已取消删除"})}))},handleEdit:function(e,t){console.log(e,t)},handleDelete:function(e,t){console.log(e,t)}},components:{}},Object(L["a"])(q,"data",(function(){return{tableData:[{ip:"127.0.0.1",port:"80",local:"北京",is_live:!0},{ip:"127.0.0.2",port:"80",local:"上海",is_live:!1},{ip:"127.0.0.1",port:"80",local:"北京",is_live:!0},{ip:"127.0.0.1",port:"80",local:"北京",is_live:!0},{ip:"127.0.0.1",port:"80",local:"北京",is_live:!0},{ip:"127.0.0.1",port:"80",local:"北京",is_live:!0},{ip:"127.0.0.1",port:"80",local:"北京",is_live:!0}],search:""}})),Object(L["a"])(q,"created",(function(){var e=this;z().then((function(t){e.tableData=t.data})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t)}))})),Object(L["a"])(q,"methods",{handleWarn:function(e){this.$router.push({path:"/warn/detail",query:{ip:e.ip,local:e.local}})},handleSend:function(e){this.$router.push({path:"/send/detail",query:{ip:e.ip,local:e.local}})},handleHistory:function(e){this.$router.push("/send/"+e.ip+"/"+e.local)},handleShow:function(e){this.$router.push("/show/"+e.ip+"/"+e.local)},formateAgent:function(e,t,n){return e.port?e.ip+":"+e.port:e.ip},formateLive:function(e,t,n){return n?"是":"否"}}),q),A=V,M=(n("cccb"),Object(u["a"])(A,$,x,!1,null,null,null)),N=M.exports,F=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"about"},[n("SingleChart")],1)},H=[],Q=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("div",{staticClass:"nav"},[n("div",{staticClass:"subNav"},[n("el-select",{staticClass:"select",attrs:{placeholder:"时长"},model:{value:e.selector,callback:function(t){e.selector=t},expression:"selector"}},e._l(e.options,(function(e){return n("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1),n("div",{staticClass:"block"},[n("span",{staticClass:"demonstration"}),n("el-date-picker",{attrs:{type:"datetimerange","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期","value-format":"yyyy-MM-dd HH:mm:ss","default-time":["08:00:00","12:00:00"],"default-value":new Date},model:{value:e.timePickerValue,callback:function(t){e.timePickerValue=t},expression:"timePickerValue"}})],1),n("el-row",[n("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.search()}}},[n("div",{staticClass:"icon"},[e._v("\n            搜索\n          ")])])],1)],1)]),n("div",{staticClass:"single-chart"},[n("div",{staticClass:"echart",attrs:{name:e.name,id:"main"}})])])},B=[],W=(n("6b54"),{name:"SingleChart",data:function(){return{options:[{value:"30",label:"30m"},{value:"60",label:"60m"},{value:"300",label:"5h"},{value:"720",label:"12h"},{value:"1440",label:"24h"}],selector:"120",timePickerValue:[],elId:null,name:1,chartTitle:null,noDataTip:!1,config:"",myChart:"",runtime:[[30,40],[40,50],[30,40],[45,60],[50,40],[30,40],[40,50],[30,40],[45,60],[80,40],[30,40],[45,60],[50,40],[30,40],[45,60],[50,40],[30,40],[45,60],[50,40]],thisTime:Date()}},computed:{cpu:function(){for(var e=[],t=0;t<this.runtime.length;t++)e.push(this.runtime[t][0]);return e},mem:function(){for(var e=[],t=0;t<this.runtime.length;t++)e.push(this.runtime[t][1]);return e}},props:["data1"],created:function(){},watch:{cpu:function(e){this.drawChart()},mem:function(e){this.drawChart()},data1:function(e,t){this.cData=e}},destroyed:function(){clearInterval(this.interval)},methods:{drawChart:function(){for(var e=[],t=new Date,a=0;a<this.cpu.length;a++){var i=new Date(t.setMinutes(t.getMinutes()-1)),l=i.getDate(),r=i.getMonth()+1;String(r).length<2&&(r="0"+r),String(l).length<2&&(l="0"+l);var o=r+"-"+l,s=i.getHours();s<10&&(s="0"+s);var c=i.getMinutes();c<10&&(c="0"+c);var u=i.getSeconds();u<10&&(u="0"+u);var d=o+"\n"+s+":"+c+":"+u;e.unshift(d)}n("3eba");var p="main",h=this.$echarts.init(document.getElementById(p)),m={title:{text:"runtime"},tooltip:{trigger:"axis"},legend:{data:["Cpu","Mem"]},grid:{left:"3%",right:"4%",bottom:"3%",containLabel:!0},toolbox:{feature:{saveAsImage:{},restore:{}}},xAxis:{type:"category",boundaryGap:!1,axisLine:{lineStyle:{color:"#a1a1a2"}},splitLine:{show:!0,lineStyle:{color:["#a1a1a2"],width:1,type:"solid"}},data:e},yAxis:[{type:"value",name:"Cpu",axisLine:{lineStyle:{color:"#a1a1a2"}},splitLine:{show:!0,lineStyle:{color:["#a1a1a2"],width:1,type:"solid"}}},{type:"value",name:"Mem",axisLabel:{formatter:"{value} %"}}],series:[{type:"line",name:"Cpu",showSymbol:!1,data:this.cpu},{type:"line",name:"Mem",showSymbol:!1,data:this.mem}]};h.setOption(m)},getUlId:function(){console.log("methods",this.elId)},search:function(){console.log(this.selector),console.log(this.value1);var e=this.value1[0].toString();console.log(e)},now:function(){var e=new Date,t=new Date(e.setMinutes(e.getMinutes()-this.selector));this.timePickerValue=[t,new Date]}},mounted:function(){this.drawChart(),this.now()},components:{}}),U=W,G=(n("356b"),Object(u["a"])(U,Q,B,!1,null,"e008381a",null)),J=G.exports,Y={name:"about",components:{SingleChart:J}},R=Y,Z=Object(u["a"])(R,F,H,!1,null,null,null),K=Z.exports,X=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[e._m(0),n("div",{staticClass:"table-container"},[n("el-table",{attrs:{data:e.tableData.filter((function(t){return!e.search||t.local.toLowerCase().includes(e.search.toLowerCase())||t.ip.toLowerCase().includes(e.search.toLowerCase())})),fit:""}},[n("el-table-column",{attrs:{label:"Host",formatter:e.formateAgent}}),n("el-table-column",{attrs:{prop:"local",label:"区域"}}),n("el-table-column",{attrs:{formatter:e.formateMetric,prop:"metric",label:"监控指标"}}),n("el-table-column",{attrs:{formatter:e.formateLive,prop:"is_live",width:"100",label:"是否存活"}}),n("el-table-column",{attrs:{align:"center"},scopedSlots:e._u([{key:"header",fn:function(t){return[n("el-input",{attrs:{size:"mini",placeholder:"输入区域或ip关键字搜索"},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})]}},{key:"default",fn:function(t){return[n("el-button",{attrs:{type:"success",size:"mini"},on:{click:function(n){return e.handleEdit(t.row)}}},[e._v("Edit")])]}}])})],1)],1)])},ee=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"headline"},[n("h3",[e._v(" 告警配置 ")])])}],te={name:"warnConfig",components:{},data:function(){return{tableData:[{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,metric:["cpu利用率","mem"]},{ip:"127.0.0.2",port:"2016-05-03",local:"上海",is_live:!1,metric:["cpu利用率","mem"]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,metric:["cpu利用率","mem"]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,metric:["cpu利用率","mem"]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,metric:["cpu利用率","mem"]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,metric:["cpu利用率","mem"]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,metric:["cpu利用率","mem"]}],search:""}},created:function(){var e=this;z().then((function(t){e.tableData=t.data})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t)}))},methods:{handleEdit:function(e){this.$router.push({path:"/warn/detail",query:{ip:e.ip,local:e.local}})},formateMetric:function(e,t,n){return n.join(" , ").toString()},formateAgent:function(e,t,n){return e.port?e.ip+":"+e.port:e.ip},formateLive:function(e,t,n){return n?"是":"否"}}},ne=te,ae=(n("7f2f"),Object(u["a"])(ne,X,ee,!1,null,null,null)),ie=ae.exports,le=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[e._m(0),n("div",{staticClass:"table-container"},[n("el-table",{attrs:{data:e.tableData.filter((function(t){return!e.search||t.local.toLowerCase().includes(e.search.toLowerCase())||t.ip.toLowerCase().includes(e.search.toLowerCase())})),fit:""}},[n("el-table-column",{attrs:{label:"Host",formatter:e.formateAgent}}),n("el-table-column",{attrs:{prop:"local",label:"区域"}}),n("el-table-column",{attrs:{formatter:e.formateSend,prop:"send",label:"配置类型"}}),n("el-table-column",{attrs:{formatter:e.formateLive,prop:"is_live",width:"100",label:"是否存活"}}),n("el-table-column",{attrs:{align:"center"},scopedSlots:e._u([{key:"header",fn:function(t){return[n("el-input",{attrs:{size:"mini",placeholder:"输入区域或ip关键字搜索"},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})]}},{key:"default",fn:function(t){return[n("el-button",{attrs:{type:"success",size:"mini"},on:{click:function(n){return e.handleEdit(t.row)}}},[e._v("Edit")])]}}])})],1)],1)])},re=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"headline"},[n("h3",[e._v(" 发送配置 ")])])}],oe={0:"email",1:"kafka",2:"nsq",3:"http"},se={email:0,kafka:1,nsq:2,http:3};function ce(e){var t=oe[e];return t||"unknown"}function ue(e){var t=oe[e];return!!t}var de={name:"sendConfig",components:{},data:function(){return{tableData:[{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,send:[1,2]},{ip:"127.0.0.2",port:"2016-05-03",local:"上海",is_live:!1,send:[1,2]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,send:[1,2]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,send:[1,2]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,send:[1,2]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,send:[1,2]},{ip:"127.0.0.1",port:"2016-05-03",local:"北京",is_live:!0,send:[1,2]}],search:""}},created:function(){var e=this;I().then((function(t){e.tableData=t.data})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t)}))},methods:{handleEdit:function(e){this.$router.push({path:"/send/detail",query:{ip:e.ip,local:e.local}})},formateSend:function(e,t,n){for(var a=[],i=0;i<n.length;i++)a.push(ce(n[i]));return a.join(" , ").toString()},formateAgent:function(e,t,n){return e.port?e.ip+":"+e.port:e.ip},formateLive:function(e,t,n){return n?"是":"否"}}},pe=de,he=(n("64ff"),Object(u["a"])(pe,le,re,!1,null,null,null)),me=he.exports,fe=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"warning"},[n("div",{staticClass:"content"},[e._m(0),n("el-table",{staticStyle:{width:"100%"},attrs:{data:e.tableData.filter((function(t){return!e.search||t.id==e.search.toLowerCase()||t.level.toLowerCase().includes(e.search.toLowerCase())||t.tabName.toLowerCase().includes(e.search.toLowerCase())||t.startTime.includes(e.search)})),"header-cell-style":{height:"100px"}}},[n("el-table-column",{attrs:{label:"对象",prop:"id",align:"center","header-aligh":"center"}}),n("el-table-column",{attrs:{label:"告警内容",prop:"tabName",align:"center","header-aligh":"center"}}),n("el-table-column",{attrs:{label:"级别",prop:"level",align:"center","header-aligh":"center"}}),n("el-table-column",{attrs:{"min-width":"100px",label:"开始时间",prop:"startTime",align:"center","header-aligh":"center"}}),n("el-table-column",{attrs:{label:"异常值",prop:"outliers",align:"center","header-aligh":"center"}}),n("el-table-column",{attrs:{label:"阈值",prop:"threshold",align:"center","header-aligh":"center"}}),n("el-table-column",{attrs:{label:"持续时间",prop:"during",align:"center","header-aligh":"center"}}),n("el-table-column",{attrs:{align:"center","min-width":"100px"},scopedSlots:e._u([{key:"header",fn:function(t){return[n("el-input",{attrs:{size:"mini",placeholder:"输入关键字进行搜索 "},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})]}},{key:"default",fn:function(t){return[n("el-button",{attrs:{size:"mini"},on:{click:function(n){return e.handleEdit(t.$index,t.row)}}},[e._v("Edit")]),n("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(n){return e.openDelete(t.$index,t.row)}}},[e._v("Delete")])]}}])})],1)],1)])},ve=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"headline"},[n("h3",[e._v("告警信息")])])}],be={name:"warning",components:{},data:function(){return{dialogVisible:!1,search:"",tableData:[{id:1,tabName:"cpu利用率",level:"严重",startTime:"2022-02-10 18:00:00",outliers:85,threshold:80,during:30},{id:2,tabName:"cpu利用率",level:"严重",startTime:"2022-02-10 19:00:00",outliers:85,threshold:80,during:30},{id:3,tabName:"mem利用率",level:"中等",startTime:"2022-02-10 20:00:00",outliers:85,threshold:80,during:30}]}},methods:{deleteData:function(e,t){this.dialogVisible=!1,console.log(e),console.log(t),this.tableData.splice(e,1)},openDelete:function(e,t){var n=this;this.$confirm("此操作将永久删除该文件, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){n.deleteData(e,t),n.$message({type:"success",message:"删除成功!"})})).catch((function(){n.$message({type:"info",message:"已取消删除"})}))},handleEdit:function(e,t){console.log(e,t)},handleDelete:function(e,t){console.log(e,t)}}},ge=be,ye=(n("ecac"),Object(u["a"])(ge,fe,ve,!1,null,null,null)),_e=ye.exports,we=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"warnListChart"},[n("PieChart")],1)},ke=[],Ce=function(){var e=this,t=e.$createElement;e._self._c;return e._m(0)},$e=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("div",{staticClass:"single-chart"},[n("div",{staticClass:"echart",attrs:{id:"main"}},[e._v("pieChart")])])])}],xe={name:"PieChart",data:function(){return{}},computed:{},props:[],created:function(){},watch:{},destroyed:function(){},methods:{run:function(e){echarts.registerTransform(window.ecSimpleTransform.aggregate),option={dataset:[{id:"raw",source:e},{id:"since_year",fromDatasetId:"raw",transform:[{type:"filter",config:{dimension:"Year",gte:1950}}]},{id:"income_aggregate",fromDatasetId:"since_year",transform:[{type:"ecSimpleTransform:aggregate",config:{resultDimensions:[{name:"min",from:"Income",method:"min"},{name:"Q1",from:"Income",method:"Q1"},{name:"median",from:"Income",method:"median"},{name:"Q3",from:"Income",method:"Q3"},{name:"max",from:"Income",method:"max"},{name:"Country",from:"Country"}],groupBy:"Country"}},{type:"sort",config:{dimension:"Q3",order:"asc"}}]}],title:{text:"Income since 1950"},tooltip:{trigger:"axis",confine:!0},xAxis:{name:"Income",nameLocation:"middle",nameGap:30,scale:!0},yAxis:{type:"category"},grid:{bottom:100},legend:{selected:{detail:!1}},dataZoom:[{type:"inside"},{type:"slider",height:20}],series:[{name:"boxplot",type:"boxplot",datasetId:"income_aggregate",itemStyle:{color:"#b8c5f2"},encode:{x:["min","Q1","median","Q3","max"],y:"Country",itemName:["Country"],tooltip:["min","Q1","median","Q3","max"]}},{name:"detail",type:"scatter",datasetId:"since_year",symbolSize:6,tooltip:{trigger:"item"},label:{show:!0,position:"top",align:"left",verticalAlign:"middle",rotate:90,fontSize:12},itemStyle:{color:"#d00000"},encode:{x:"Income",y:"Country",label:"Year",itemName:"Year",tooltip:["Country","Year","Income"]}}]},myChart.setOption(option)}},mounted:function(){},components:{}},Le=xe,Te=(n("2649"),Object(u["a"])(Le,Ce,$e,!1,null,"407bdbe1",null)),De=Te.exports,Se={name:"warnListChart",components:{PieChart:De}},Ee=Se,je=Object(u["a"])(Ee,we,ke,!1,null,null,null),Oe=je.exports,Pe=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"container"},[n("div",{staticClass:"headline"},[n("h3",[e._v(" "+e._s(e.ip)+" - "+e._s(e.local)+" 判定配置详情 ")])]),n("div",{staticClass:"table-container"},[n("el-row",[n("el-col",{attrs:{span:18}},[n("div",[e._v(" ")])]),n("el-col",{attrs:{span:6}},[n("div",[n("el-input",{attrs:{size:"mini",placeholder:"输入关键字搜索"},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})],1)])],1),n("el-table",{attrs:{data:e.tableData.filter((function(t){return!e.search||t.metric.toLowerCase().includes(e.search.toLowerCase())||t.period.toLowerCase().includes(e.search.toLowerCase())||t.threshold.toLowerCase().includes(e.search.toLowerCase())})),fit:""}},[n("el-table-column",{attrs:{label:"指标类型",prop:"metric"}}),n("el-table-column",{attrs:{formatter:e.formateMethod,prop:"method",label:"聚合方式"}}),n("el-table-column",{attrs:{prop:"period",label:"聚合周期"}}),n("el-table-column",{attrs:{align:"center",label:"管理"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(n){return e.showEdit(t.row)}}},[e._v("编辑配置")]),-1!=t.row.id?[n("el-button",{attrs:{type:"danger",size:"mini"},on:{click:function(n){return e.handleDel(t.row)}}},[e._v("删除配置")])]:[n("el-button",{attrs:{type:"info",disabled:"",size:"mini"}},[e._v("默认配置")])]]}}])})],1)],1),n("el-dialog",{attrs:{title:"配置编辑",visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[n("el-form",[n("el-form-item",{attrs:{required:""}},[n("el-col",{attrs:{span:3}},[e._v("主机:")]),n("el-col",{attrs:{span:9}},[n("el-input",{attrs:{disabled:""},model:{value:e.ip,callback:function(t){e.ip=t},expression:"ip"}})],1),n("el-col",{attrs:{span:3}},[e._v("区域:")]),n("el-col",{attrs:{span:9}},[n("el-input",{attrs:{disabled:""},model:{value:e.local,callback:function(t){e.local=t},expression:"local"}})],1)],1),n("el-form-item",{attrs:{required:""}},[n("el-col",{attrs:{span:3}},[e._v("指标类型:")]),n("el-col",{attrs:{span:9}},[n("el-input",{attrs:{disabled:""},model:{value:e.metric,callback:function(t){e.metric=t},expression:"metric"}})],1),n("el-col",{attrs:{span:3}},[e._v("聚合周期:")]),n("el-col",{attrs:{span:9}},[n("el-input",{model:{value:e.period,callback:function(t){e.period=t},expression:"period"}})],1)],1),n("el-form-item",{attrs:{required:""}},[n("el-col",{attrs:{span:3}},[e._v("聚合方式:")]),n("el-col",{attrs:{span:21}},[n("el-select",{attrs:{placeholder:"请选择聚合方式"},model:{value:e.method,callback:function(t){e.method=t},expression:"method"}},e._l(e.methodList,(function(e,t){return n("el-option",{key:t,attrs:{label:e,value:t}})})),1)],1)],1),e._l(e.threshold,(function(t,a){return n("el-form-item",{key:a},[n("el-col",{attrs:{span:3}},[e._v(e._s(t.label)+"阈值:")]),n("el-col",{attrs:{span:21}},[n("el-input",{model:{value:t.value,callback:function(n){e.$set(t,"value",n)},expression:"item.value"}})],1)],1)})),n("p",{staticClass:"tip"},[e._v("设置告警阈值, 留空则不监控")])],2),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:function(t){e.dialogFormVisible=!1}}},[e._v("取 消")]),n("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.handleEdit()}}},[e._v("确 定")])],1)],1)],1)},qe=[],ze={0:"总和",1:"平均值",2:"中位数",3:"积分",4:"极值",5:"标准差",6:"最大值",7:"最小值"},Ie={"总和":0,"平均值":1,"中位数":2,"积分":3,"极值":4,"标准差":5,"最大值":6,"最小值":7};function Ve(e){var t=ze[e];return t||"unknown"}function Ae(e){var t=ze[e];return!!t}n("456d"),n("ac6a");function Me(e){return JSON.stringify(e)}function Ne(e){return JSON.parse(e)}var Fe={0:"info",1:"warn",2:"error",3:"panic"},He={info:0,warn:1,error:2,panic:3};function Qe(e){var t=Ne(e),n={};return Object.keys(Fe).forEach((function(e){n[Fe[e]]={label:Fe[e],value:t[e]?t[e]:""}})),n}function Be(e){var t={};return Object.keys(e).forEach((function(n){e[n].value&&(t[He[n]]=e[n].value)})),Me(t)}function We(e){var t=!1;return Object.keys(Fe).forEach((function(n){e[Fe[n]].value&&(t=!0)})),t}function Ue(e){var t=Fe[e];return t||"unknown"}function Ge(e){var t=Fe[e];return!!t}function Je(e,t){var n="/judgment/info/"+e+"/"+t;return P(n,{})}function Ye(e,t,n,a,i,l,r){var o="/judgment/update/"+e+"/"+t+"/"+n;return O(o,{method:a,metric:i,period:l,threshold:r})}function Re(e,t,n){var a="/judgment/del/"+e+"/"+t+"/"+n;return O(a,{})}var Ze={name:"warnDetail",components:{},data:function(){return{id:-1,ip:"",local:"",metric:"",method:1,period:"5m",threshold:{},dialogFormVisible:!1,tableData:[{id:-1,ip:"127.0.0.1",local:"北京",metric:"cpu利用率",method:1,period:"5m",threshold:'{"0":0.01, "1": 0.02}'},{id:11,ip:"127.0.0.1",local:"广州",metric:"mem",method:2,period:"5m",threshold:"{}"},{id:12,ip:"127.0.0.1",local:"深圳",metric:"rate",method:3,period:"5m",threshold:"{}"},{id:13,ip:"127.0.0.1",local:"上海",metric:"cpu利用率",method:0,period:"5m",threshold:'{"1":0.01, "2": 0.02}'},{id:14,ip:"127.0.0.1",local:"北京",metric:"cpu利用率",method:1,period:"5m",threshold:'{"3":0.01}'}],search:"",methodList:ze}},created:function(){var e=this;this.ip=this.$route.query.ip,this.local=this.$route.query.local,this.ip&&this.local||this.$router.push({path:"/"}),this.metric="",this.id=-1,this.method=1,this.period="5m",this.threshold={},Je(this.ip,this.local).then((function(t){e.tableData=t.data})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t)}))},methods:{check:function(){return this.ip&&this.local&&this.metric&&this.period?Ae(this.method)?!!We(this.threshold)||(this.$alert("至少需要设置一个告警阈值, 请重试~"),!1):(this.$alert("聚合方式有误, 请重试~"),!1):(this.$alert("操作有误, 请重试~"),!1)},showEdit:function(e){this.metric=e.metric,this.method=Ve(e.method),this.period=e.period,this.id=e.id,this.threshold=Qe(e.threshold),this.dialogFormVisible=!0},handleEdit:function(){var e=this;Ae(this.method)||(this.method=Ie[this.method]),this.check()&&Ye(this.ip,this.local,this.id,this.method,this.metric,this.period,Be(this.threshold)).then((function(t){e.$alert("更新成功"),e.$router.go(0)})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t),e.$router.go(0)}))},handleDel:function(e){var t=this;Re(this.ip,this.local,e.id).then((function(e){t.$alert("删除成功"),t.$router.go(0)})).catch((function(e){e.msg?t.$alert(e.msg):t.$alert(e),t.$router.go(0)}))},formateMethod:function(e,t,n){return Ve(n)}}},Ke=Ze,Xe=(n("8767"),Object(u["a"])(Ke,Pe,qe,!1,null,null,null)),et=Xe.exports,tt=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"container"},[n("div",{staticClass:"headline"},[n("h3",[e._v(" "+e._s(e.ip)+" - "+e._s(e.local)+" 发送配置详情 ")])]),n("div",{staticClass:"table-container"},[n("el-row",[n("el-col",{attrs:{span:1}},[n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(t){return e.showAdd()}}},[e._v("新增配置")])],1),n("el-col",{attrs:{span:17}},[n("div",[e._v(" ")])]),n("el-col",{attrs:{span:6}},[n("div",[n("el-input",{attrs:{size:"mini",placeholder:"输入关键字搜索"},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})],1)])],1),n("el-table",{attrs:{data:e.tableData.filter((function(t){return!e.search||t.ip.toLowerCase().includes(e.search.toLowerCase())||t.local.toLowerCase().includes(e.search.toLowerCase())||t.config.toLowerCase().includes(e.search.toLowerCase())})),fit:""}},[n("el-table-column",{attrs:{formatter:e.formateLevel,label:"告警等级",prop:"level"}}),n("el-table-column",{attrs:{formatter:e.formateSendType,prop:"sendType",label:"告警类型"}}),n("el-table-column",{attrs:{align:"center",label:"管理"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{type:"primary",size:"mini"},on:{click:function(n){return e.showEdit(t.row)}}},[e._v("编辑配置")]),-1!=t.row.id?[n("el-button",{attrs:{type:"danger",size:"mini"},on:{click:function(n){return e.handleDel(t.row)}}},[e._v("删除配置")])]:[n("el-button",{attrs:{type:"info",disabled:"",size:"mini"}},[e._v("默认配置")])]]}}])})],1)],1),n("el-dialog",{attrs:{title:"配置编辑",visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[n("el-form",[n("el-form-item",{attrs:{required:""}},[n("el-col",{attrs:{span:3}},[e._v("主机:")]),n("el-col",{attrs:{span:9}},[n("el-input",{attrs:{disabled:""},model:{value:e.ip,callback:function(t){e.ip=t},expression:"ip"}})],1),n("el-col",{attrs:{span:3}},[e._v("区域:")]),n("el-col",{attrs:{span:9}},[n("el-input",{attrs:{disabled:""},model:{value:e.local,callback:function(t){e.local=t},expression:"local"}})],1)],1),n("el-form-item",{attrs:{required:""}},[n("el-col",{attrs:{span:3}},[e._v("告警等级:")]),n("el-col",{attrs:{span:9}},[n("el-select",{attrs:{placeholder:"请选择告警等级"},model:{value:e.level,callback:function(t){e.level=t},expression:"level"}},e._l(e.levelList,(function(e,t){return n("el-option",{key:t,attrs:{label:e,value:t}})})),1)],1),n("el-col",{attrs:{span:3}},[e._v("告警类型:")]),n("el-col",{attrs:{span:9}},[n("el-select",{attrs:{placeholder:"请选择告警类型"},on:{change:e.TypeChange},model:{value:e.sendType,callback:function(t){e.sendType=t},expression:"sendType"}},e._l(e.sendList,(function(e,t){return n("el-option",{key:t,attrs:{label:e,value:t}})})),1)],1)],1),e._l(e.config,(function(t,a){return n("el-form-item",{key:a},[n("el-col",{attrs:{span:3}},[e._v(e._s(t.label)+":")]),n("el-col",{attrs:{span:21}},[t.selectList?n("el-select",{attrs:{placeholder:"请选择"},model:{value:t.value,callback:function(n){e.$set(t,"value",n)},expression:"item.value"}},e._l(t.selectList,(function(e){return n("el-option",{key:e,attrs:{label:e,value:e}})})),1):n("el-input",{model:{value:t.value,callback:function(n){e.$set(t,"value",n)},expression:"item.value"}})],1),t.tip?n("p",{staticClass:"tip"},[e._v(e._s(t.tip))]):e._e()],1)}))],2),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:function(t){e.dialogFormVisible=!1}}},[e._v("取 消")]),n("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.handleEdit()}}},[e._v("确 定")])],1)],1)],1)},nt=[];function at(e,t){var n="/send/info/"+e+"/"+t;return P(n,{})}function it(e,t,n,a,i){console.log(e,t,n,a,i);var l="/send/info/"+e+"/"+t;return O(l,{sendType:n,level:a,config:i})}function lt(e,t,n,a,i,l){console.log(e,t,n,a,i,l);var r="/send/update/"+e+"/"+t+"/"+n;return O(r,{sendType:a,level:i,config:l})}function rt(e,t,n){var a="/send/del/"+e+"/"+t+"/"+n;return O(a,{})}n("f559");var ot={0:{target:{label:"邮箱地址",value:"",tip:"多个邮箱则以,分割"},format_type:{label:"告警格式",value:"line",selectList:["line","json","html"]}},1:{topic:{label:"Topic",value:""},format_type:{label:"告警格式",value:"line",selectList:["line","json","html"]},address:{label:"集群地址",value:"",tip:"多个地址则以,分割"},version:{label:"版本号",value:"",tip:"kafka版本,如0.2.0.0"},partition_type:{label:"选区方式",value:"random",selectList:["random","robin","hash","manual"]},partition:{label:"分区号",value:"",tip:"分区号, 选区方式为manual时必填"},partition_key:{label:"分区hash",value:"",tip:"分区号, 选区方式为hash时必填"}},2:{topic:{label:"Topic",value:""},format_type:{label:"告警格式",value:"line",selectList:["line","json","html"]},address:{label:"nsq地址",value:""}},3:{url:{label:"目标URL",value:""},format_type:{label:"告警格式",value:"line",selectList:["line","json","html"]},method:{label:"请求方式",value:"POST",selectList:["GET","POST","HEAD","PUT","DELETE","CONNECT","OPTIONS","TRACE"]}}};function st(e,t){var n=Ne(t),a={},i=ot[e];return Object.keys(i).forEach((function(e){a[e]={label:i[e].label,value:n[e]?n[e]:i[e].value},i[e].selectList&&(a[e].selectList=i[e].selectList),i[e].tip&&(a[e].tip=i[e].tip)})),a}function ct(e){var t={};return Object.keys(e).forEach((function(n){t[n]=e[n].value})),Me(t)}function ut(e,t){if(0==e){if(!t.target.value)return{check:!1,msg:"邮箱地址不能为空"};if(-1==t.target.value.indexOf("@"))return{check:!1,msg:"邮箱地址格式有误"}}if(1==e){if(!t.address.value)return{check:!1,msg:"地址不能为空"};if(!t.topic.value)return{check:!1,msg:"topic不能为空"};if("manual"==t.partition_type.value&&!t.partition.value)return{check:!1,msg:"选区方式为manual时,分区号不能为空"};if("hash"==t.partition_type.value&&!t.partition_key.value)return{check:!1,msg:"选区方式为hash时,分区hash值不能为空"}}if(2==e){if(!t.topic.value)return{check:!1,msg:"topic不能为空"};if(!t.address.value)return{check:!1,msg:"地址不能为空"}}if(3==e){if(!t.url.value)return{check:!1,msg:"请求地址不能为空"};if(t.url.value.startsWith("http://")||t.url.value.startsWith("https://"))return{check:!1,msg:"请求地址有误"}}return{check:!0,msg:""}}var dt={name:"sendDetail",components:{},data:function(){return{id:-1,ip:"",local:"",sendType:"",level:"",config:{},sendList:oe,levelList:Fe,search:"",dialogFormVisible:!1,tableData:[{id:1,ip:"127.0.0.1",local:"北京",sendType:0,level:0,config:'{"target": "526756656@qq.com,123456789@163.com", "format_type": "html"}'},{id:11,ip:"127.0.0.1",local:"广州",sendType:1,level:1,config:'{"topic": "test", "format_type": "line", "address": "127.0.0.1:8554,127.0.0.2:7894", "version": "0.2.0.1", "partition_type": "hash","partition":0, "partition_key": "test"}'},{id:12,ip:"127.0.0.1",local:"深圳",sendType:2,level:2,config:'{"topic": "test_nsq", "format_type": "json", "address": "127.0.0.1:1234"}'},{id:13,ip:"127.0.0.1",local:"上海",sendType:0,level:3,config:'{"target": "526756656@qq.com", "format_type": "line"}'},{id:14,ip:"127.0.0.1",local:"北京",sendType:3,level:0,config:'{"url": "http://127.0.0.1/test", "format_type": "json", "method": "GET"}'}]}},created:function(){var e=this;this.ip=this.$route.query.ip,this.local=this.$route.query.local,this.ip&&this.local||this.$router.push({path:"/"}),this.id=-1,this.sendType=ce(1),this.level=Ue(1),this.config={},at(this.ip,this.local).then((function(t){e.tableData=t.data})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t)}))},methods:{check:function(){if(!this.ip||!this.local||!this.config)return this.$alert("操作有误, 请重试~"),!1;if(!ue(this.sendType))return this.$alert("告警类型有误, 请重试~"),!1;if(!Ge(this.level))return this.$alert("告警类型有误, 请重试~"),!1;var e=ut(this.sendType,this.config);return!!e.check||(this.$alert(e.msg),!1)},formateSendType:function(e,t,n){return ce(n)},formateLevel:function(e,t,n){return Ue(n)},TypeChange:function(e){this.config=ot[e]},showAdd:function(){this.level=Ue(0),this.sendType=ce(0),this.id=-1,this.config=ot[0],this.dialogFormVisible=!0},showEdit:function(e){this.level=Ue(e.level),this.sendType=ce(e.sendType),this.id=e.id,this.config=st(e.sendType,e.config),this.dialogFormVisible=!0},handleEdit:function(){var e=this;ue(this.sendType)||(this.sendType=se[this.sendType]),Ge(this.level)||(this.level=He[this.level]),this.check()&&(-1==this.id?it(this.ip,this.local,this.sendType,this.level,ct(this.config)).then((function(t){e.$router.go(0)})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t),e.$router.go(0)})):lt(this.ip,this.local,this.id,this.sendType,this.level,ct(this.config)).then((function(t){e.$router.go(0)})).catch((function(t){t.msg?e.$alert(t.msg):e.$alert(t),e.$router.go(0)})))},handleDel:function(e){var t=this;rt(this.ip,this.local,e.id).then((function(e){t.$alert("删除成功"),t.$router.go(0)})).catch((function(e){e.msg?t.$alert(e.msg):t.$alert(e),t.$router.go(0)}))}}},pt=dt,ht=(n("8013"),Object(u["a"])(pt,tt,nt,!1,null,null,null)),mt=ht.exports,ft=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"about"},[n("div",{staticClass:"headline"},[n("h3",[e._v(" "+e._s(e.ip)+"-"+e._s(e.local)+" ")])]),n("SingleChart")],1)},vt=[],bt={name:"show",created:function(){this.ip=this.$route.params.ip,this.local=this.$route.params.local},components:{SingleChart:J}},gt=bt,yt=Object(u["a"])(gt,ft,vt,!1,null,null,null),_t=yt.exports;a["default"].use(C["a"]);var wt=new C["a"]({mode:"history",base:"/",routes:[{path:"/",name:"home",component:N},{path:"/about",name:"about",component:K},{path:"/warn",name:"warnConfig",component:ie},{path:"/send",name:"sendConfig",component:me},{path:"/warning",name:"warning",component:_e},{path:"/warnListChart",name:"warnListChart",component:Oe},{path:"/warn/detail",name:"warnDetail",component:et},{path:"/send/detail",name:"sendDetail",component:mt},{path:"/show/:ip/:local",name:"showDetail",component:_t}]}),kt=n("2f62");a["default"].use(kt["a"]);var Ct=new kt["a"].Store({state:{},mutations:{},actions:{}}),$t=n("9483");Object($t["a"])("".concat("/","service-worker.js"),{ready:function(){console.log("App is being served from cache by a service worker.\nFor more details, visit https://goo.gl/AFskqB")},registered:function(){console.log("Service worker has been registered.")},cached:function(){console.log("Content has been cached for offline use.")},updatefound:function(){console.log("New content is downloading.")},updated:function(){console.log("New content is available; please refresh.")},offline:function(){console.log("No internet connection found. App is running in offline mode.")},error:function(e){console.error("Error during service worker registration:",e)}});var xt=n("313e"),Lt=n("5c96"),Tt=n.n(Lt);n("c69f");a["default"].use(Tt.a);var Dt=n("1f94"),St=n.n(Dt);a["default"].prototype.$axios=D.a,a["default"].prototype.qs=j.a,a["default"].use(St.a),a["default"].config.productionTip=!1,a["default"].prototype.$echarts=xt,new a["default"]({router:wt,store:Ct,render:function(e){return e(k)}}).$mount("#app")},5845:function(e,t,n){},"5ac2":function(e,t,n){},"5c48":function(e,t,n){},6343:function(e,t,n){},"64ff":function(e,t,n){"use strict";n("81c0")},"7c55":function(e,t,n){"use strict";n("5c48")},"7f2f":function(e,t,n){"use strict";n("5845")},8013:function(e,t,n){"use strict";n("f3ef")},"81c0":function(e,t,n){},8767:function(e,t,n){"use strict";n("ba42")},a9f6:function(e,t,n){"use strict";n("6343")},ba42:function(e,t,n){},c69f:function(e,t,n){},cccb:function(e,t,n){"use strict";n("d563")},d563:function(e,t,n){},e7c6:function(e,t,n){},ecac:function(e,t,n){"use strict";n("5ac2")},f3ef:function(e,t,n){}});
//# sourceMappingURL=app.aba386ea.js.map