// 该文件专门用于创建整个应用的路由
import VueRouter from "vue-router";

import Login from "@/pages/Login";
import Upload from "@/pages/Upload";
import axios from "axios";

const router = new VueRouter({
    routes:[
        {
            path:"/login",
            component:Login
        },
        {
            path:"/upload",
            component:Upload
        }
    ]
});

// 全局前置路由守卫，在页面初始化和路由跳转之前被调用
router.beforeEach((to,from,next) => {
    // console.log("前置路由守卫被触发")
    // console.log("from：");
    // console.log(from)
    // console.log("to：");
    // console.log(to);

    if(to.path.startsWith("/login")){
        localStorage.removeItem("x-token");
        next()
    }else{
        let x_token = localStorage.getItem("x-token");
        if (x_token !== null){
            // 有token，验证token
            axios.get(
                "/api/checkToken",
                {
                    params:{
                        "token":x_token
                    }
                }
            ).then(resp =>{
                // console.log(resp)
                if (resp.data.message === "OK"){
                    next()
                }else {
                    next("/login")
                }
            }).catch(error => {
                console.log(error)
                next("/login")
            })
        }
    }


    next();  // 放行请求
})



export default router