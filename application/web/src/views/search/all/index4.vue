<template>
  <div>
    <input type="file" @change="handleFileUpload">
    <button @click="uploadImage">上传图片</button>
    <img :src="imageUrl" alt="图片">
  </div>
</template>

<script>
import { uploadimg, getimg } from '@/api/account';

export default {
  data() {
    return {
      imageUrl: '',
      file: null
    };
  },
  mounted() {
    this.getImage();
  },
  methods: {
    getImage() {
        getimg()
        .then(response => {
         // const blob = this.base64ToBlob(respones.data);
          var blob = new Blob([response.data], {type: 'image/png'});
          alert(response)
          var imageUrl = URL.createObjectURL(blob);
         // this.imageUrl = response.data;
          alert(response.data)
        })
        .catch(error => {
          alert(error.message)
          console.log(error);
        });
    },
    handleFileUpload(event) {
      this.file = event.target.files[0];
    },
    uploadImage() {
      const formData = new FormData();
      formData.append('image', this.file);
      uploadimg(formData)
      .then(response => {
        console.log(response.data)
      })
      .catch(error => {
        console.log(error)
      })
    }
  }
};
</script>
