import {RouteRecordRaw} from "vue-router";

const routes: RouteRecordRaw[] = [
	{
		path: "/index",
		name: "home",
		component: () => import("@/views/index/index.vue"),
	},
	{
		path:"/recommend",
		name:"recommend",
		component: () => import("@/views/Recommend/index.vue"),
	},
	{
		path:"/mine",
		name:"mine",
		component: () => import("@/views/Mine/index.vue"),
	},
	{
		path:"/err/404",
		name:"err404",
		component:()=>import("@/views/404/index.vue"),
	}
]
export default routes;