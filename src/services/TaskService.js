import axios from 'axios'

export default {
  getTasks() {
    return axios.get('/api/tasks')
  },
  
  getTask(id) {
    return axios.get(`/api/tasks/${id}`)
  },
  
  createTask(task) {
    return axios.post('/api/tasks', task)
  },
  
  updateTask(task) {
    return axios.put(`/api/tasks/${task.id}`, task)
  },
  
  deleteTask(id) {
    return axios.delete(`/api/tasks/${id}`)
  }
}
