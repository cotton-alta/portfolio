<template>
  
</template>

<script>
export default {
  mounted() {
    this.$store.dispatch("setUser", null)
    this.$router.push("/edit/login")
  }
}
</script>