<template>
  <div>
    <h1>麻将祝计算</h1>

    <form @submit.prevent="calculate">
      <div>
        <h3>胡牌玩家</h3>
        <label>玩家ID：</label>
        <input v-model="request.player_id" required /><br />

        <label>胡牌类型：</label>
        <select v-model="request.hu_type" required>
          <option value="ron">荣胡</option>
          <option value="tsumo">自摸</option>
        </select><br />

        <label v-if="request.hu_type === 'ron'">被荣胡玩家ID：</label>
        <input v-if="request.hu_type === 'ron'" v-model="request.target_id" required /><br />

        <label>里宝数量：</label>
        <input v-model.number="request.li_bao_count" type="number" min="0" required /><br />

        <label>赤宝数量：</label>
        <input v-model.number="request.chi_bao_count" type="number" min="0" required /><br />

        <label>是否役满：</label>
        <input type="checkbox" v-model="request.is_yakuman" /><br />
      </div>

      <button type="submit">计算祝</button>
    </form>

    <h2>玩家祝点数</h2>
    <ul v-if="scores">
      <li v-for="(score, player) in scores" :key="player">
        {{ player }}: {{ score }} 祝
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      request: {
        player_id: '',
        hu_type: 'ron',
        target_id: '',
        li_bao_count: 0,
        chi_bao_count: 0,
        is_yakuman: false,
      },
      scores: null,  // 存储返回的玩家祝点数
    };
  },
  methods: {
    async calculate() {
      try {
        const response = await axios.post('http://localhost:8081/calculate', this.request);
        this.scores = response.data.scores;  // 更新所有玩家的祝点数
      } catch (error) {
        console.error('计算失败', error);
      }
    },
  },
};
</script>
