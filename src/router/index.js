import { createRouter, createWebHistory } from 'vue-router'
import TaskList from '@/views/TaskList.vue'
import TaskForm from '@/views/TaskForm.vue'

const routes = [
  {
    path: '/',
    name: 'TaskList',
    component: TaskList
  },
  {
    path: '/create',
    name: 'CreateTask',
    component: TaskForm
  },
  {
    path: '/edit/:id',
    name: 'EditTask',
    component: TaskForm,
    props: true
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
