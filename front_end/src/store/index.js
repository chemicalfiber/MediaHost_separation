// 该文件用于创建Vuex中最核心的store

// 引入Vue
import Vue from "vue"
// 引入Vuex
import Vuex from "vuex"
Vue.use(Vuex)   // 在创建store对象之前必须先使用该插件，否则会报错！

// 准备actions——用于响应组件中的动作
const actions = {
    // actions中的函数名一般小写
    // context是迷你版的store对象，value是调用此action时传递的参数
    jia(context,value){
        console.log("action中：",context,value)
        // 通过context调用mutations中名为"JIA"的函数，将value传过去
        context.commit("JIA",value)
    },
    jian(context,value){
        // 通过context调用mutations中名为"JIAN"的函数，将value传过去
        context.commit("JIAN",value)
    },
    jiaOdd(context,value){
        if (context.state.sum % 2 === 0){
            // console.log("现在和是偶数，不加")
            context.dispatch("notOddNotice",value); // 将这个action转发到另一个action处理
        }else {
            context.commit("JIA",value) // 通过context调用mutations中名为"JIA"的函数，将value传过去
        }
    },
    jiaWait(context,value){
        setTimeout(()=>{
            context.commit("JIA",value)
        },1000);
    },
    notOddNotice(context,value){
        console.log("现在和是偶数："+context.state.sum+"，不加")
    }
}
// 准备mutations——用于操作数据
const mutations = {
    // mutations中的函数名一般大写
    // state是存储数据的地方，value是调用context.commit()时传递的值
    JIA(state, value){
        console.log("mutations中：",state,value)
        // 操作state
        state.sum += value;
    },
    JIAN(state, value){
        state.sum -= value;
    }
}
// 准备state——用于存放数据
const state = {
    sum:0,   // 在state身上定义初始数据，用于被其他组件共享
    username: "mike"
}

// 创建store
const store = new Vuex.Store({
    // 传入三个配置项
    actions,
    mutations,
    state
});
// 导出store对象，让其他js可以引入
export default store