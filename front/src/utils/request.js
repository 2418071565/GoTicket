import axios from 'axios'

// 创建 axios 实例
const request = axios.create({
  timeout: 5000, // 请求超时时间
  headers: {
    // 移除默认的 Content-Type，让浏览器自动设置
  }
})

// 请求拦截器
request.interceptors.request.use(
  config => {
    // 从 localStorage 获取 token
    const token = localStorage.getItem('token')
    if (token) {
      config.headers['Authorization'] = token
    }
    return config
  },
  error => {
    console.error('Request error:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  response => response,
  error => {
    const { response } = error;
    
    // 处理 401 未授权错误
    if (response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
      return Promise.reject(new Error('Please login first'));
    }

    // 处理后端返回的错误信息
    if (response?.data?.msg) {
      return Promise.reject(new Error(response.data.msg));
    }

    return Promise.reject(error);
  }
)

// 封装 GET 请求
export const get = (url, params) => {
  return request({
    method: 'get',
    url,
    params
  })
}

// 封装 POST 请求
export const post = (url, data) => {
    return request({
        method: 'post',
        url,
        data,
        headers: {
          // 如果是 FormData，不要手动设置 Content-Type
          ...(!(data instanceof FormData) && {
            'Content-Type': 'application/json'
          })
        }
      });
}

// 封装 PUT 请求
export const put = (url, data) => {
  return request({
    method: 'put',
    url,
    data
  })
}

// 封装 DELETE 请求
export const del = (url, params) => {
  return request({
    method: 'delete',
    url,
    params
  })
}

export default request