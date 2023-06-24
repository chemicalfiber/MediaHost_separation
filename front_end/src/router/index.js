// 该文件专门用于创建整个应用的路由
import VueRouter from "vue-router";

import Login from "@/pages/Login";
import Upload from "@/pages/Upload";
import Register from "@/pages/Rgister"
import axios from "axios";
import Video from "@/pages/video/Videos";
import Images from "@/pages/image/Images";
import Me from "@/pages/Me";
import Doc from "@/pages/Doc";
import VideoDetail from "@/pages/video/VideoDetail";
import ImageDetail from "@/pages/image/ImageDetail";

const router = new VueRouter({
    routes: [
        {
            path: "/register",
            component: Register,
            meta:{
                title:"注册"
            }
        },
        {
            path: "/login",
            component: Login,
            meta:{
                title:"登录"
            }
        },
        {
            path: "/",
            component: Upload,
            meta:{
                title:"上传"
            }
        },
        {
            path:"/videos",
            component:Video,
            meta:{
                title:"视频列表"
            }
        },
        {path:"/videos/:id",component:VideoDetail,meta:{title:"视频详情"}},
        {
            path:"/images",
            component:Images,
            meta:{
                title:"图片列表"
            }
        },
        {path:"/images/:id",component:ImageDetail,meta:{title:"图片详情"}},
        {
            path:"/me",
            component:Me,
            meta:{
                title:"个人中心"
            }
        },{
            path:"/doc",
            component:Doc,
            meta:{
                title:"使用文档"
            }
        }
    ]
});

// 全局前置路由守卫，在页面初始化和路由跳转之前被调用
router.beforeEach((to, from, next) => {
    // console.log("前置路由守卫被触发")
    // console.log("from：");
    // console.log(from)
    // console.log("to：");
    // console.log(to);

    if(to.meta.title){
        document.title = to.meta.title;
    }
    if (to.path.startsWith("/login") || to.path.startsWith("/register")) {
        next()
    } else {
        let x_token = localStorage.getItem("x-token");
        if (x_token !== null) {
            // 有token，验证token
            axios.get(
                "/api/checkToken",
                {
                    params: {
                        "x-token": x_token
                    }
                }
            ).then(resp => {
                // console.log(resp)
                if (resp.data.message === "OK") {
                    next()
                } else {
                    next("/login")
                }
            }).catch(error => {
                console.log(error)
                next("/login")
            })
        }
    }
})

export default router