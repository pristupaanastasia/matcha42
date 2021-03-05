<template>
  <div>
    <form @submit.prevent="checkForm">
      <p v-if="errors.length">
        <b>Пожалуйста исправьте указанные ошибки:</b>
      <ul>
         <li class="list-group-item" v-for="error in errors"  v-bind:key="error"></li>
      </ul>
      </p>
      <p>
        <label for="login">Login</label>
        <input
            id="login"
            v-model="login"
            type="text"
            name="login"
        >
      </p>
      <p>
      <label for="email">Email</label>
      <input
          id="email"
          v-model="email"
          type="text"
          name="email"
      >
    </p>
    <p>
      <label for="password">Password</label>
      <input
          id="password"
          v-model="password"
          type="text"
          name="password"
      >
    </p>
    <p>
      <label for="first_name">First Name</label>
      <input
          id="first_name"
          v-model="first_name"
          type="text"
          name="first_name"
      >
      </p>
    <p>
      <label for="last_name">Last Name</label>
      <input
          id="last_name"
          v-model="last_name"
          type="text"
          name="last_name"
      >
      </p>
      <input type="button" value="Отправить" onclick="submit()">

    </form>
  </div>
</template>
<script>

    const apiUrl = 'https://localhost:9000/register'


    export default {
        el: '#app',
        data() {
          return {
            errors: [],
            login: null,
            email: null,
            password: null,
            first_name: null,
            last_name: null,
          }
        },


        methods: {
            checkForm: function () {
                this.errors = [];
                if (this.email && this.login && this.password && this.first_name && this.last_name) {

                    return true;
                }


                if (!this.email) {
                    this.errors.push('Требуется указать email.');
                }else if (!this.validEmail(this.email)) {
                    this.errors.push('Укажите корректный адрес электронной почты.');
                }
                if (!this.password) {
                    this.errors.push('Требуется указать пароль.');
                }
                if (!this.first_name) {
                    this.errors.push('Требуется указать имя.');
                }
                if (!this.last_name) {
                    this.errors.push('Требуется указать фамилия.');
                }
                if (!this.login) {
                    this.errors.push('Требуется указать логин.');
                }


            },
          async submit() {
            let response = await fetch(apiUrl);
            if (response.ok) {
              fetch(apiUrl,
              {
                method: 'POST',
                    headers: {
                      'Content-Type': 'form-data'
              },
                body: this.data(),
                redirect: 'follow'
              });
            }else{
              alert("Ошибка HTTP: " + response.status);
            }
          }
        },


    };
</script>