var app = new Vue({
  el: '.main',
  data: {
    showChannelsBlock: true,
    showPlayerBlock: false,
    tg: window.Telegram.WebApp,
    channel_name: '',
  },
  methods: {

    onImageClick(channel_name) {
      this.showPlayerBlock = true;
      this.showChannelsBlock = false;
      this.channel_name = channel_name;
      this.$nextTick(function () {
        this.autoplay(channel_name);
      });
    },

    autoplay(channel_name) {
      if (Hls.isSupported()) {
        var hls = new Hls();
        hls.loadSource('https://ucdn.beetv.kz/btv/live/hls/000004499.m3u8');
        hls.attachMedia(this.$refs.video);
        this.$refs.video.play();
      } else if (this.$refs.video.canPlayType('application/vnd.apple.mpegurl')) {
        this.$refs.video.src = 'https://ucdn.beetv.kz/btv/live/hls/000004499.m3u8';
        this.$refs.video.addEventListener('loadedmetadata', function () {
          this.$refs.video.play();
        }.bind(this));
      }
    }, 

    close_webapp() {
      var context = {
        status: "success"
      }

      this.tg.sendData(JSON.stringify(context));
      this.tg.close();
    }
  },
});
