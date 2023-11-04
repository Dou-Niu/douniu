import { RouteRecordRaw } from "vue-router";

const HomeRoutes: RouteRecordRaw[] = [
	{
		path: "/index",
		name: "首页",
		component: () => import("@/views/index/index.vue"),
	},
	{
		path: "/recommend",
		name: "推荐",
		component: () => import("@/views/play/index.vue"),
	},
	{
		path: "/follow",
		name: "关注",
		component: () => import("@/views/play/index.vue"),
	},
	{
		path: "/mine",
		name: "我的",
		component: () => import("@/views/mine/index.vue"),
	},
	{
		path: "/upload",
		name: "upload",
		component: () => import("@/views/upload/index.vue"),
		meta: {
			dontShow: true,
		}
	},
	{
		path: "/play",
		name: "play",
		component: () => import("@/views/play/index.vue"),
		meta: {
			dontShow: true,
		}
	},
]
//1-游戏 2-生活 3-影视 4-动漫 5-知识
const ChannelRoutes: RouteRecordRaw[] = [
	{
		path: "/search/:id",
		name: "search",
		component: () => import("@/views/search/index.vue"),
		meta: {
			dontShow: true,
		}
	},
	{
		path: "/channel/:id",
		name: "channel",
		component: () => import("@/views/channel/index.vue"),
		meta: {
			dontShow: true,
		}
	},
]
const ErrRoutes: RouteRecordRaw[] = [
	{
		path: "/err/404",
		name: "404",
		component: () => import("@/views/404/index.vue"),
	}
]
export { HomeRoutes, ErrRoutes, ChannelRoutes};