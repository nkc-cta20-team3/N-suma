/**
 * plugins/vuetify/index.js
 * included in `../index.js`
 */

import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

const NsumaLightTheme = {
  dark: false,
  colors: {
    background: "#FFFFFF",
    surface: "#FFFFFF",
    primary: "#ff9800",
    secondary: "#ffeb3b",
    accent: "#cddc39",
    error: "#e91e63",
    warning: "#f44336",
    info: "#00bcd4",
    success: "#009688",
  },
};

/*
// 背景色のみ変えた試作
const NsumaDarkTheme = {
  dark: true,
  colors: {
    background: '#121212',
    surface: '#121212',
    primary: #ff9800,
    secondary: #ffeb3b,
    accent: #cddc39,
    error: #e91e63,
    warning: #f44336,
    info: #00bcd4,
    success: #009688
    }
}
*/

/*
  // デフォルトのテーマカラー 
  colors: {
    background: '#FFFFFF',
    surface: '#FFFFFF',
    primary: '#6200EE',
    secondary: '#03DAC6',
    error: '#B00020',
    info: '#2196F3',
    success: '#4CAF50',
    warning: '#FB8C00',
  },
  // サンプルのテーマカラー
  colors: {
    background: '#FFFFFF',
    surface: '#FFFFFF',
    primary: '#4caf50',
    'primary-darken-1': '#3700B3',
    secondary: '#8bc34a',
    'secondary-darken-1': '#018786',
    error: '#ffeb3b',
    info: '#ff5722',
    success: '#795548',
    warning: '#ffc107',
    accent: "#cddc39",
  },
*/

export default createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: "NsumaLightTheme",
    dark: false,
    themes: {
      NsumaLightTheme,
    },
  },
});
