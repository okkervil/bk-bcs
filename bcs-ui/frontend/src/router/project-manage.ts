// 项目管理
const OperateAudit = () => import(/* webpackChunkName: 'project' */'@/views/mc/operate-audit.vue');
const EventQuery = () => import(/* webpackChunkName: 'project' */'@/views/mc/new-event-query.vue');
const tencentCloud = () => import(/* webpackChunkName: 'project' */'@/views/cloudtoken/tencentCloud.vue');

export default [
  {
    path: 'operate-audit',
    component: OperateAudit,
    name: 'operateAudit',
  },
  {
    path: 'event-query',
    name: 'eventQuery',
    component: EventQuery,
  },
  {
    path: 'tencent-cloud',
    name: 'tencentCloud',
    component: tencentCloud,
    meta: {
      title: 'Tencent Cloud',
      hideBack: true,
    },
  },
];