<template>
  <div class="container">
    <tabTitle :pageTitle="pageTitle"/>
    <div class="logo-wrapper">
      <skill class="logo-card" :skill="skill" v-for="skill in skills" :key="skill.id"/>
    </div>
    <!-- /.logo-wrapper -->
  </div>
  <!-- /.container -->
</template>

<script>
import tabTitle from "~/components/layouts/tabTitle.vue"
import skill from "~/components/ui/skill.vue"
import apache from "~/static/logo/apache.png"
import centos from "~/static/logo/centos.png"
import css from "~/static/logo/css.png"
import django from "~/static/logo/django.png"
import html from "~/static/logo/html.png"
import javascript from "~/static/logo/javascript.png"
import mongo from "~/static/logo/mongo.png"
import node from "~/static/logo/node.png"
import nuxt from "~/static/logo/nuxt.png"
import python from "~/static/logo/python.png"
import vue from "~/static/logo/vue.png"
import git from "~/static/logo/git.png"
import dynamo from "~/static/logo/dynamo.png"
import s3 from "~/static/logo/s3.png"
import element from "~/static/logo/element.png"
import go from "~/static/logo/go.png"

export default {
  components: {
    tabTitle,
    skill
  },
  data() {
    return {
      pageTitle: "SKILLS",
      skills: {
        html: { 
          name: "HTML5", 
          img: html,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9733;"
        },
        css: {
          name: "CSS3",
          img: css,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9733;"
        },
        javascript: {
          name: "JavaScript",
          img: javascript, 
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9733;"
        },
        node: {
          name: "Node.js",
          img: node,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9734;"
        },
        vue: {
          name: "Vue.js",
          img: vue,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9733;"
        },
        nuxt: {
          name: "Nuxt.js",
          img: nuxt,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9734;" 
        },
        Python: {
          name: "Python",
          img: python,
          years: "1年",
          level: "&#9733;&#9733;&#9733;&#9734;&#9734;"  
        },
        django: {
          name: "Django",
          img: django,
          years: "1年",
          level: "&#9733;&#9733;&#9734;&#9734;&#9734;"
        },
        go: {
          name: "Go",
          img: go,
          years: "1年",
          level: "&#9733;&#9733;&#9734;&#9734;&#9734;" 
        },
        centos: {
          name: "CentOS",
          img: centos,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9734;&#9734;"
        },
        apache: {
          name: "Apache",
          img: apache,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9734;"
        },
        mongo: {
          name: "MongoDB",
          img: mongo,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9734;"
        },
        DynamoDB: {
          name: "DynamoDB",
          img: dynamo,
          years: "1年",
          level: "&#9733;&#9733;&#9734;&#9734;&#9734;"
          },
        S3: { 
          name: "S3",
          img: s3,
          years: "1年",
          level: "&#9733;&#9733;&#9733;&#9734;&#9734;"
        },
        Element: {
          name: "Element UI",
          img: element,
          years: "1年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9734;"
          },
        git: {
          name: "Git",
          img: git,
          years: "2年",
          level: "&#9733;&#9733;&#9733;&#9733;&#9733;"  
        },
      }
    }
  }
}
</script>

<style lang="scss" scoped>

.container {
  display: block;
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  min-height: calc(100vh - 50px);
  @include mq {
    min-height: calc(100vh - 80px);
  }
}

.logo {
  &-wrapper {
    margin: 0 auto;
    margin-bottom: 30px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    flex-wrap: wrap;
    &::after{
      content:"";
      display: block;
      width:30%;
    }
    @include mq {
      width: 600px;
      justify-content: space-between;
    }
  }
  &-card {
    margin: 20px 10px;
    @include mq(sm) {
      margin: 20px 30px;
    }
  }
}

</style>