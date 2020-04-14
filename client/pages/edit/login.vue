<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="form-wrapper">
      <label for="username" class="form-label">ユーザID</label>
      <div class="form-text">
        <input type="text" name="username" v-model="username" />
      </div>
      <!-- /.form-text -->
      <label for="password" class="form-label">パスワード</label>
      <div class="form-text">
        <input type="password" name="password" v-model="password" />
      </div>
      <!-- /.form-text -->
      <div class="form-submit"><input @click="login" type="submit" value="ログイン"/></div>
    </div>
    <!-- /.form-wrapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import axios from "axios"

export default {
  middleware({ store, app, redirect }) {
  },
  components: {
    tabTitle
  },
  data() {
    return { 
      pageTitle: "LOGIN",
      username: "",
      password: ""
    }
  },
  methods: {
    login: function() {
      console.log(this.username, this.password)
      this.$axios.$get("api/login", {
        auth: {
          username: this.username,
          password: this.password
        }
      })
        .then(result => {
          this.$store.dispatch("setUser", result)
          this.$router.push("/edit/createWork")
        })
        .catch(err => {
          console.log("err", err)
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  display: block;
  min-height: calc(100vh - 80px);
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
}

.form {
  &-wrapper {
    display: flex;
    width: 100%;
    flex-direction: column;
    align-items: center;
  }
  &-label {
    max-width: 480px;
    width: 75%;
    text-align: left;
    font-size: calc(15px + 0.3vw);
    font-weight: bolder;
    color: $menu-background;
    margin-top: 20px;
  }
  &-text {
    max-width: 500px;
    width: 80%;
    margin: 10px 0px;
    input {
      width: 100%;
      font-size: calc(15px + 0.3vw);
      line-height: 2rem;
      border-radius: 5px;
      display: inline-block;
    }
  }
  &-submit {
    width: 60%;
    max-width: 300px;
    input {
      font-size: calc(15px + 0.3vw);
      font-weight: bolder;
      background-color: $menu-background;
      color: $main-background;
      width: 100%;
      margin-top: 40px;
      line-height: 2rem;
      border-radius: 5px;
      display: inline-block;
    }
  }
}
</style>