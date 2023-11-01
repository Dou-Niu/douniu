import {RouteRecordRaw} from "vue-router";

const HomeRoutes: RouteRecordRaw[] = [
	{
		path: "/index",
		name: "首页",
		component: () => import("@/views/index/index.vue"),
	},
	{
		path:"/recommend",
		name:"推荐",
		component: () => import("@/views/recommend/index.vue"),
	},
	{
		path:"/mine",
		name:"我的",
		component: () => import("@/views/mine/index.vue"),
	},
]
const ErrRoutes: RouteRecordRaw[] = [
	{
		path:"/err/404",
		name:"404",
		component: () => import("@/views/404/index.vue"),
	}
]
export {HomeRoutes,ErrRoutes};