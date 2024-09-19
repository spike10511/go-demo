<template>
  <div>
    <h2>Register</h2>
    <form @submit.prevent="register">
      <div>
        <label for="username">Username:</label>
        <input v-model="username" type="text" id="username" required>
      </div>
      <div>
        <label for="password">Password:</label>
        <input v-model="password" type="password" id="password" required>
      </div>
      <button type="submit">Register</button>
    </form>
    <p v-if="error">{{ error }}</p>
  </div>
</template>

<script>
export default {
  name: 'UserRegister',  // 更新组件名
  data() {
    return {
      username: '',
      password: '',
      error: ''
    }
  },
  methods: {
    async register() {
      try {
        const response = await fetch('http://localhost:8080/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            username: this.username,
            password: this.password
          })
        });
        const data = await response.json();
        if (response.ok) {
          alert('Registration successful');
        } else {
          this.error = data.error;
        }
      } catch (err) {
        this.error = 'Failed to register';
      }
    }
  }
}
</script>
