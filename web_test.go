package main

import (
	"net/http"
	"testing"
)

func TestWeb(t *testing.T) {
	f := NewDanmaku("abc")
	f.WriteDanmaku([]byte(`<html lang="en" class="fac-events-icons-ready fac-events-typefaces-cera-pro-ready fac-events-typefaces-cera-pro-700-normal-ready fac-events-typefaces-cera-pro-400-normal-ready fac-events-typefaces-cera-pro-400-italic-ready fac-events-typefaces-cera-pro-700-italic-ready fac-events-typefaces-cera-round-pro-ready fac-events-typefaces-cera-round-pro-700-normal-ready fac-events-typefaces-cera-round-pro-400-normal-ready fac-events-typefaces-cera-round-pro-900-normal-ready fac-events-typefaces-cera-round-pro-500-normal-ready fa5-events-icons-ready fa5-events-typefaces-fa5-proxima-nova-ready fa5-events-typefaces-fa5-proxima-nova-normal-normal-ready fa5-events-typefaces-fa5-proxima-nova-600-normal-ready fa5-events-typefaces-fa5-proxima-nova-bold-normal-ready fa5-events-typefaces-fa5-proxima-nova-normal-italic-ready"><head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="csrf-token" content="Ky9NECgnKEghfTcJEy4Za34NHCwNJF1dRwuJxdNypLCqQ_410cumatjm">
<style type="text/css">:root, :host {
  --fa-font-solid: normal 900 1em/1 "Font Awesome 6 Solid";
  --fa-font-regular: normal 400 1em/1 "Font Awesome 6 Regular";
  --fa-font-light: normal 300 1em/1 "Font Awesome 6 Light";
  --fa-font-thin: normal 100 1em/1 "Font Awesome 6 Thin";
  --fa-font-duotone: normal 900 1em/1 "Font Awesome 6 Duotone";
  --fa-font-brands: normal 400 1em/1 "Font Awesome 6 Brands";
}

svg:not(:root).svg-inline--fa, svg:not(:host).svg-inline--fa {
  overflow: visible;
  box-sizing: content-box;
}

.svg-inline--fa {
  display: var(--fa-display, inline-block);
  height: 1em;
  overflow: visible;
  vertical-align: -0.125em;
}
.svg-inline--fa.fa-2xs {
  vertical-align: 0.1em;
}
.svg-inline--fa.fa-xs {
  vertical-align: 0em;
}
.svg-inline--fa.fa-sm {
  vertical-align: -0.0714285705em;
}
.svg-inline--fa.fa-lg {
  vertical-align: -0.2em;
}
.svg-inline--fa.fa-xl {
  vertical-align: -0.25em;
}
.svg-inline--fa.fa-2xl {
  vertical-align: -0.3125em;
}
.svg-inline--fa.fa-pull-left {
  margin-right: var(--fa-pull-margin, 0.3em);
  width: auto;
}
.svg-inline--fa.fa-pull-right {
  margin-left: var(--fa-pull-margin, 0.3em);
  width: auto;
}
.svg-inline--fa.fa-li {
  width: var(--fa-li-width, 2em);
  top: 0.25em;
}
.svg-inline--fa.fa-fw {
  width: var(--fa-fw-width, 1.25em);
}

.fa-layers svg.svg-inline--fa {
  bottom: 0;
  left: 0;
  margin: auto;
  position: absolute;
  right: 0;
  top: 0;
}

.fa-layers-counter, .fa-layers-text {
  display: inline-block;
  position: absolute;
  text-align: center;
}

.fa-layers {
  display: inline-block;
  height: 1em;
  position: relative;
  text-align: center;
  vertical-align: -0.125em;
  width: 1em;
}
.fa-layers svg.svg-inline--fa {
  -webkit-transform-origin: center center;
          transform-origin: center center;
}

.fa-layers-text {
  left: 50%;
  top: 50%;
  -webkit-transform: translate(-50%, -50%);
          transform: translate(-50%, -50%);
  -webkit-transform-origin: center center;
          transform-origin: center center;
}

.fa-layers-counter {
  background-color: var(--fa-counter-background-color, #ff253a);
  border-radius: var(--fa-counter-border-radius, 1em);
  box-sizing: border-box;
  color: var(--fa-inverse, #fff);
  line-height: var(--fa-counter-line-height, 1);
  max-width: var(--fa-counter-max-width, 5em);
  min-width: var(--fa-counter-min-width, 1.5em);
  overflow: hidden;
  padding: var(--fa-counter-padding, 0.25em 0.5em);
  right: var(--fa-right, 0);
  text-overflow: ellipsis;
  top: var(--fa-top, 0);
  -webkit-transform: scale(var(--fa-counter-scale, 0.25));
          transform: scale(var(--fa-counter-scale, 0.25));
  -webkit-transform-origin: top right;
          transform-origin: top right;
}

.fa-layers-bottom-right {
  bottom: var(--fa-bottom, 0);
  right: var(--fa-right, 0);
  top: auto;
  -webkit-transform: scale(var(--fa-layers-scale, 0.25));
          transform: scale(var(--fa-layers-scale, 0.25));
  -webkit-transform-origin: bottom right;
          transform-origin: bottom right;
}

.fa-layers-bottom-left {
  bottom: var(--fa-bottom, 0);
  left: var(--fa-left, 0);
  right: auto;
  top: auto;
  -webkit-transform: scale(var(--fa-layers-scale, 0.25));
          transform: scale(var(--fa-layers-scale, 0.25));
  -webkit-transform-origin: bottom left;
          transform-origin: bottom left;
}

.fa-layers-top-right {
  top: var(--fa-top, 0);
  right: var(--fa-right, 0);
  -webkit-transform: scale(var(--fa-layers-scale, 0.25));
          transform: scale(var(--fa-layers-scale, 0.25));
  -webkit-transform-origin: top right;
          transform-origin: top right;
}

.fa-layers-top-left {
  left: var(--fa-left, 0);
  right: auto;
  top: var(--fa-top, 0);
  -webkit-transform: scale(var(--fa-layers-scale, 0.25));
          transform: scale(var(--fa-layers-scale, 0.25));
  -webkit-transform-origin: top left;
          transform-origin: top left;
}

.fa-1x {
  font-size: 1em;
}

.fa-2x {
  font-size: 2em;
}

.fa-3x {
  font-size: 3em;
}

.fa-4x {
  font-size: 4em;
}

.fa-5x {
  font-size: 5em;
}

.fa-6x {
  font-size: 6em;
}

.fa-7x {
  font-size: 7em;
}

.fa-8x {
  font-size: 8em;
}

.fa-9x {
  font-size: 9em;
}

.fa-10x {
  font-size: 10em;
}

.fa-2xs {
  font-size: 0.625em;
  line-height: 0.1em;
  vertical-align: 0.225em;
}

.fa-xs {
  font-size: 0.75em;
  line-height: 0.0833333337em;
  vertical-align: 0.125em;
}

.fa-sm {
  font-size: 0.875em;
  line-height: 0.0714285718em;
  vertical-align: 0.0535714295em;
}

.fa-lg {
  font-size: 1.25em;
  line-height: 0.05em;
  vertical-align: -0.075em;
}

.fa-xl {
  font-size: 1.5em;
  line-height: 0.0416666682em;
  vertical-align: -0.125em;
}

.fa-2xl {
  font-size: 2em;
  line-height: 0.03125em;
  vertical-align: -0.1875em;
}

.fa-fw {
  text-align: center;
  width: 1.25em;
}

.fa-ul {
  list-style-type: none;
  margin-left: var(--fa-li-margin, 2.5em);
  padding-left: 0;
}
.fa-ul > li {
  position: relative;
}

.fa-li {
  left: calc(var(--fa-li-width, 2em) * -1);
  position: absolute;
  text-align: center;
  width: var(--fa-li-width, 2em);
  line-height: inherit;
}

.fa-border {
  border-color: var(--fa-border-color, #eee);
  border-radius: var(--fa-border-radius, 0.1em);
  border-style: var(--fa-border-style, solid);
  border-width: var(--fa-border-width, 0.08em);
  padding: var(--fa-border-padding, 0.2em 0.25em 0.15em);
}

.fa-pull-left {
  float: left;
  margin-right: var(--fa-pull-margin, 0.3em);
}

.fa-pull-right {
  float: right;
  margin-left: var(--fa-pull-margin, 0.3em);
}

.fa-beat {
  -webkit-animation-name: fa-beat;
          animation-name: fa-beat;
  -webkit-animation-delay: var(--fa-animation-delay, 0);
          animation-delay: var(--fa-animation-delay, 0);
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 1s);
          animation-duration: var(--fa-animation-duration, 1s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, ease-in-out);
          animation-timing-function: var(--fa-animation-timing, ease-in-out);
}

.fa-bounce {
  -webkit-animation-name: fa-bounce;
          animation-name: fa-bounce;
  -webkit-animation-delay: var(--fa-animation-delay, 0);
          animation-delay: var(--fa-animation-delay, 0);
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 1s);
          animation-duration: var(--fa-animation-duration, 1s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, cubic-bezier(0.28, 0.84, 0.42, 1));
          animation-timing-function: var(--fa-animation-timing, cubic-bezier(0.28, 0.84, 0.42, 1));
}

.fa-fade {
  -webkit-animation-name: fa-fade;
          animation-name: fa-fade;
  -webkit-animation-delay: var(--fa-animation-delay, 0);
          animation-delay: var(--fa-animation-delay, 0);
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 1s);
          animation-duration: var(--fa-animation-duration, 1s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, cubic-bezier(0.4, 0, 0.6, 1));
          animation-timing-function: var(--fa-animation-timing, cubic-bezier(0.4, 0, 0.6, 1));
}

.fa-beat-fade {
  -webkit-animation-name: fa-beat-fade;
          animation-name: fa-beat-fade;
  -webkit-animation-delay: var(--fa-animation-delay, 0);
          animation-delay: var(--fa-animation-delay, 0);
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 1s);
          animation-duration: var(--fa-animation-duration, 1s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, cubic-bezier(0.4, 0, 0.6, 1));
          animation-timing-function: var(--fa-animation-timing, cubic-bezier(0.4, 0, 0.6, 1));
}

.fa-flip {
  -webkit-animation-name: fa-flip;
          animation-name: fa-flip;
  -webkit-animation-delay: var(--fa-animation-delay, 0);
          animation-delay: var(--fa-animation-delay, 0);
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 1s);
          animation-duration: var(--fa-animation-duration, 1s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, ease-in-out);
          animation-timing-function: var(--fa-animation-timing, ease-in-out);
}

.fa-shake {
  -webkit-animation-name: fa-shake;
          animation-name: fa-shake;
  -webkit-animation-delay: var(--fa-animation-delay, 0);
          animation-delay: var(--fa-animation-delay, 0);
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 1s);
          animation-duration: var(--fa-animation-duration, 1s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, linear);
          animation-timing-function: var(--fa-animation-timing, linear);
}

.fa-spin {
  -webkit-animation-name: fa-spin;
          animation-name: fa-spin;
  -webkit-animation-delay: var(--fa-animation-delay, 0);
          animation-delay: var(--fa-animation-delay, 0);
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 2s);
          animation-duration: var(--fa-animation-duration, 2s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, linear);
          animation-timing-function: var(--fa-animation-timing, linear);
}

.fa-spin-reverse {
  --fa-animation-direction: reverse;
}

.fa-pulse,
.fa-spin-pulse {
  -webkit-animation-name: fa-spin;
          animation-name: fa-spin;
  -webkit-animation-direction: var(--fa-animation-direction, normal);
          animation-direction: var(--fa-animation-direction, normal);
  -webkit-animation-duration: var(--fa-animation-duration, 1s);
          animation-duration: var(--fa-animation-duration, 1s);
  -webkit-animation-iteration-count: var(--fa-animation-iteration-count, infinite);
          animation-iteration-count: var(--fa-animation-iteration-count, infinite);
  -webkit-animation-timing-function: var(--fa-animation-timing, steps(8));
          animation-timing-function: var(--fa-animation-timing, steps(8));
}

@media (prefers-reduced-motion: reduce) {
  .fa-beat,
.fa-bounce,
.fa-fade,
.fa-beat-fade,
.fa-flip,
.fa-pulse,
.fa-shake,
.fa-spin,
.fa-spin-pulse {
    -webkit-animation-delay: -1ms;
            animation-delay: -1ms;
    -webkit-animation-duration: 1ms;
            animation-duration: 1ms;
    -webkit-animation-iteration-count: 1;
            animation-iteration-count: 1;
    transition-delay: 0s;
    transition-duration: 0s;
  }
}
@-webkit-keyframes fa-beat {
  0%, 90% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  45% {
    -webkit-transform: scale(var(--fa-beat-scale, 1.25));
            transform: scale(var(--fa-beat-scale, 1.25));
  }
}
@keyframes fa-beat {
  0%, 90% {
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  45% {
    -webkit-transform: scale(var(--fa-beat-scale, 1.25));
            transform: scale(var(--fa-beat-scale, 1.25));
  }
}
@-webkit-keyframes fa-bounce {
  0% {
    -webkit-transform: scale(1, 1) translateY(0);
            transform: scale(1, 1) translateY(0);
  }
  10% {
    -webkit-transform: scale(var(--fa-bounce-start-scale-x, 1.1), var(--fa-bounce-start-scale-y, 0.9)) translateY(0);
            transform: scale(var(--fa-bounce-start-scale-x, 1.1), var(--fa-bounce-start-scale-y, 0.9)) translateY(0);
  }
  30% {
    -webkit-transform: scale(var(--fa-bounce-jump-scale-x, 0.9), var(--fa-bounce-jump-scale-y, 1.1)) translateY(var(--fa-bounce-height, -0.5em));
            transform: scale(var(--fa-bounce-jump-scale-x, 0.9), var(--fa-bounce-jump-scale-y, 1.1)) translateY(var(--fa-bounce-height, -0.5em));
  }
  50% {
    -webkit-transform: scale(var(--fa-bounce-land-scale-x, 1.05), var(--fa-bounce-land-scale-y, 0.95)) translateY(0);
            transform: scale(var(--fa-bounce-land-scale-x, 1.05), var(--fa-bounce-land-scale-y, 0.95)) translateY(0);
  }
  57% {
    -webkit-transform: scale(1, 1) translateY(var(--fa-bounce-rebound, -0.125em));
            transform: scale(1, 1) translateY(var(--fa-bounce-rebound, -0.125em));
  }
  64% {
    -webkit-transform: scale(1, 1) translateY(0);
            transform: scale(1, 1) translateY(0);
  }
  100% {
    -webkit-transform: scale(1, 1) translateY(0);
            transform: scale(1, 1) translateY(0);
  }
}
@keyframes fa-bounce {
  0% {
    -webkit-transform: scale(1, 1) translateY(0);
            transform: scale(1, 1) translateY(0);
  }
  10% {
    -webkit-transform: scale(var(--fa-bounce-start-scale-x, 1.1), var(--fa-bounce-start-scale-y, 0.9)) translateY(0);
            transform: scale(var(--fa-bounce-start-scale-x, 1.1), var(--fa-bounce-start-scale-y, 0.9)) translateY(0);
  }
  30% {
    -webkit-transform: scale(var(--fa-bounce-jump-scale-x, 0.9), var(--fa-bounce-jump-scale-y, 1.1)) translateY(var(--fa-bounce-height, -0.5em));
            transform: scale(var(--fa-bounce-jump-scale-x, 0.9), var(--fa-bounce-jump-scale-y, 1.1)) translateY(var(--fa-bounce-height, -0.5em));
  }
  50% {
    -webkit-transform: scale(var(--fa-bounce-land-scale-x, 1.05), var(--fa-bounce-land-scale-y, 0.95)) translateY(0);
            transform: scale(var(--fa-bounce-land-scale-x, 1.05), var(--fa-bounce-land-scale-y, 0.95)) translateY(0);
  }
  57% {
    -webkit-transform: scale(1, 1) translateY(var(--fa-bounce-rebound, -0.125em));
            transform: scale(1, 1) translateY(var(--fa-bounce-rebound, -0.125em));
  }
  64% {
    -webkit-transform: scale(1, 1) translateY(0);
            transform: scale(1, 1) translateY(0);
  }
  100% {
    -webkit-transform: scale(1, 1) translateY(0);
            transform: scale(1, 1) translateY(0);
  }
}
@-webkit-keyframes fa-fade {
  50% {
    opacity: var(--fa-fade-opacity, 0.4);
  }
}
@keyframes fa-fade {
  50% {
    opacity: var(--fa-fade-opacity, 0.4);
  }
}
@-webkit-keyframes fa-beat-fade {
  0%, 100% {
    opacity: var(--fa-beat-fade-opacity, 0.4);
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  50% {
    opacity: 1;
    -webkit-transform: scale(var(--fa-beat-fade-scale, 1.125));
            transform: scale(var(--fa-beat-fade-scale, 1.125));
  }
}
@keyframes fa-beat-fade {
  0%, 100% {
    opacity: var(--fa-beat-fade-opacity, 0.4);
    -webkit-transform: scale(1);
            transform: scale(1);
  }
  50% {
    opacity: 1;
    -webkit-transform: scale(var(--fa-beat-fade-scale, 1.125));
            transform: scale(var(--fa-beat-fade-scale, 1.125));
  }
}
@-webkit-keyframes fa-flip {
  50% {
    -webkit-transform: rotate3d(var(--fa-flip-x, 0), var(--fa-flip-y, 1), var(--fa-flip-z, 0), var(--fa-flip-angle, -180deg));
            transform: rotate3d(var(--fa-flip-x, 0), var(--fa-flip-y, 1), var(--fa-flip-z, 0), var(--fa-flip-angle, -180deg));
  }
}
@keyframes fa-flip {
  50% {
    -webkit-transform: rotate3d(var(--fa-flip-x, 0), var(--fa-flip-y, 1), var(--fa-flip-z, 0), var(--fa-flip-angle, -180deg));
            transform: rotate3d(var(--fa-flip-x, 0), var(--fa-flip-y, 1), var(--fa-flip-z, 0), var(--fa-flip-angle, -180deg));
  }
}
@-webkit-keyframes fa-shake {
  0% {
    -webkit-transform: rotate(-15deg);
            transform: rotate(-15deg);
  }
  4% {
    -webkit-transform: rotate(15deg);
            transform: rotate(15deg);
  }
  8%, 24% {
    -webkit-transform: rotate(-18deg);
            transform: rotate(-18deg);
  }
  12%, 28% {
    -webkit-transform: rotate(18deg);
            transform: rotate(18deg);
  }
  16% {
    -webkit-transform: rotate(-22deg);
            transform: rotate(-22deg);
  }
  20% {
    -webkit-transform: rotate(22deg);
            transform: rotate(22deg);
  }
  32% {
    -webkit-transform: rotate(-12deg);
            transform: rotate(-12deg);
  }
  36% {
    -webkit-transform: rotate(12deg);
            transform: rotate(12deg);
  }
  40%, 100% {
    -webkit-transform: rotate(0deg);
            transform: rotate(0deg);
  }
}
@keyframes fa-shake {
  0% {
    -webkit-transform: rotate(-15deg);
            transform: rotate(-15deg);
  }
  4% {
    -webkit-transform: rotate(15deg);
            transform: rotate(15deg);
  }
  8%, 24% {
    -webkit-transform: rotate(-18deg);
            transform: rotate(-18deg);
  }
  12%, 28% {
    -webkit-transform: rotate(18deg);
            transform: rotate(18deg);
  }
  16% {
    -webkit-transform: rotate(-22deg);
            transform: rotate(-22deg);
  }
  20% {
    -webkit-transform: rotate(22deg);
            transform: rotate(22deg);
  }
  32% {
    -webkit-transform: rotate(-12deg);
            transform: rotate(-12deg);
  }
  36% {
    -webkit-transform: rotate(12deg);
            transform: rotate(12deg);
  }
  40%, 100% {
    -webkit-transform: rotate(0deg);
            transform: rotate(0deg);
  }
}
@-webkit-keyframes fa-spin {
  0% {
    -webkit-transform: rotate(0deg);
            transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
            transform: rotate(360deg);
  }
}
@keyframes fa-spin {
  0% {
    -webkit-transform: rotate(0deg);
            transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
            transform: rotate(360deg);
  }
}
.fa-rotate-90 {
  -webkit-transform: rotate(90deg);
          transform: rotate(90deg);
}

.fa-rotate-180 {
  -webkit-transform: rotate(180deg);
          transform: rotate(180deg);
}

.fa-rotate-270 {
  -webkit-transform: rotate(270deg);
          transform: rotate(270deg);
}

.fa-flip-horizontal {
  -webkit-transform: scale(-1, 1);
          transform: scale(-1, 1);
}

.fa-flip-vertical {
  -webkit-transform: scale(1, -1);
          transform: scale(1, -1);
}

.fa-flip-both,
.fa-flip-horizontal.fa-flip-vertical {
  -webkit-transform: scale(-1, -1);
          transform: scale(-1, -1);
}

.fa-rotate-by {
  -webkit-transform: rotate(var(--fa-rotate-angle, none));
          transform: rotate(var(--fa-rotate-angle, none));
}

.fa-stack {
  display: inline-block;
  vertical-align: middle;
  height: 2em;
  position: relative;
  width: 2.5em;
}

.fa-stack-1x,
.fa-stack-2x {
  bottom: 0;
  left: 0;
  margin: auto;
  position: absolute;
  right: 0;
  top: 0;
  z-index: var(--fa-stack-z-index, auto);
}

.svg-inline--fa.fa-stack-1x {
  height: 1em;
  width: 1.25em;
}
.svg-inline--fa.fa-stack-2x {
  height: 2em;
  width: 2.5em;
}

.fa-inverse {
  color: var(--fa-inverse, #fff);
}

.sr-only,
.fa-sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border-width: 0;
}

.sr-only-focusable:not(:focus),
.fa-sr-only-focusable:not(:focus) {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border-width: 0;
}

.svg-inline--fa .fa-primary {
  fill: var(--fa-primary-color, currentColor);
  opacity: var(--fa-primary-opacity, 1);
}

.svg-inline--fa .fa-secondary {
  fill: var(--fa-secondary-color, currentColor);
  opacity: var(--fa-secondary-opacity, 0.4);
}

.svg-inline--fa.fa-swap-opacity .fa-primary {
  opacity: var(--fa-secondary-opacity, 0.4);
}

.svg-inline--fa.fa-swap-opacity .fa-secondary {
  opacity: var(--fa-primary-opacity, 1);
}

.svg-inline--fa mask .fa-primary,
.svg-inline--fa mask .fa-secondary {
  fill: black;
}

.fad.fa-inverse,
.fa-duotone.fa-inverse {
  color: var(--fa-inverse, #fff);
}</style><link rel="icon" href="/favicon.ico">
<link rel="icon" href="/images/favicon/icon.svg" type="image/svg+xml">
<link rel="apple-touch-icon" href="/images/favicon/apple-touch-icon.png">
<link rel="manifest" href="/manifest/app.json">

<meta name="theme-color" content="#528DD7">


<title>Icons | Font Awesome</title>
<meta id="meta-application-name" name="application-name" content="Icons | Font Awesome">
<meta id="meta-description" name="description" content="Used by millions of designers, devs, &amp; content creators. Open-source. Always free. Check out the new Thin style, now available in Font Awesome 6.">
<meta id="meta-keywords" name="keywords" content="icons,vector icons,svg icons,free icons,icon font,webfont,desktop icons,svg,font awesome,font awesome free,font awesome pro">

<!-- Schema.org markup for Google+ -->
<meta id="meta-item-name" itemprop="name" content="Icons | Font Awesome">
<meta id="meta-item-description" itemprop="description" content="Used by millions of designers, devs, &amp; content creators. Open-source. Always free. Check out the new Thin style, now available in Font Awesome 6.">
<meta id="meta-item-image" itemprop="image" content="https://img.fortawesome.com/349cfdf6/fontawesome-open-graph.png">

<!-- Twitter Card data -->
<meta id="twt-card" name="twitter:card" content="summary">
<meta id="twt-site" name="twitter:site" content="@fontawesome">
<meta id="twt-title" name="twitter:title" content="Icons | Font Awesome">
<meta id="twt-description" name="twitter:description" content="Used by millions of designers, devs, &amp; content creators. Open-source. Always free. Check out the new Thin style, now available in Font Awesome 6.">
<meta id="twt-creator" name="twitter:creator" content="@fontawesome">
<meta id="twt-image" name="twitter:image" content="https://img.fortawesome.com/349cfdf6/fontawesome-open-graph.png">

<!-- Open Graph data -->
<meta id="og-title" property="og:title" content="Icons | Font Awesome">
<meta id="og-type" property="og:type" content="website">
<meta id="og-url" property="og:url" content="https://fontawesome.com">
<meta id="og-image" property="og:image" content="https://img.fortawesome.com/349cfdf6/fontawesome-open-graph.png">
<meta id="og-description" property="og:description" content="Used by millions of designers, devs, &amp; content creators. Open-source. Always free. Check out the new Thin style, now available in Font Awesome 6.">


    <link rel="stylesheet" data-purpose="Layout StyleSheet" title="Default" disabled="" href="/css/app-e23924a0edb4b571d42fe2a3ca61ab8b.css?vsn=d">
    <link rel="stylesheet" data-purpose="Layout StyleSheet" title="Web Awesome" href="/css/app-wa-6972bf07d6b9f99336a9a78bb51ecbc8.css?vsn=d">
    <link rel="stylesheet" href="https://site-assets.fontawesome.com/releases/v6.1.1/css/all.css">
<script type="text/javascript" src="https://beacon-v2.helpscout.net/static/js/vendor.2901613d.js"></script><script type="text/javascript" src="https://beacon-v2.helpscout.net/static/js/main.fd07e5f9.js"></script><script type="text/javascript" async="" src="https://beacon-v2.helpscout.net"></script><script type="text/javascript" async="" src="https://www.gstatic.com/recaptcha/releases/nEGwmCAyCoKVn9PSwAGnQWhY/recaptcha__zh_cn.js" crossorigin="anonymous" integrity="sha384-IV40M14FUqdPALi8j1Bj740rlbWLj8vzwMkRNoCneIXLXdtypRKGhmPs/y4b3jsP"></script><script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-30136587-4"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'UA-30136587-4', { cookie_flags: 'max-age=7200;secure;samesite=none' });
</script>


      <script src="https://use.fortawesome.com/1ce05b4b.js"></script><link rel="stylesheet" href="https://fonticons-free-fonticons.netdna-ssl.com/kits/1ce05b4b/publications/119222/woff2.css" media="all">

    <link rel="preload" href="/js/settings-fb840945c8c0933b31d2bde8715a7317.js?vsn=d" as="script">
    <link rel="preload" href="/js/app-3b3c6c72cc6dd0ad25cb325680d3780b.js?vsn=d" as="script">

    <script defer="" src="/js/settings-fb840945c8c0933b31d2bde8715a7317.js?vsn=d"></script>
    <script defer="" src="/js/app-3b3c6c72cc6dd0ad25cb325680d3780b.js?vsn=d"></script>
<script type="text/javascript">!function(e,t,n){function a(){var e=t.getElementsByTagName("script")[0],n=t.createElement("script");n.type="text/javascript",n.async=!0,n.src="https://beacon-v2.helpscout.net",e.parentNode.insertBefore(n,e)}if(e.Beacon=n=function(t,n,a){e.Beacon.readyQueue.push({method:t,options:n,data:a})},n.readyQueue=[],"complete"===t.readyState)return a();e.attachEvent?e.attachEvent("onload",a):e.addEventListener("load",a,!1)}(window,document,window.Beacon||function(){});</script>
<script type="text/javascript">
  window.Beacon('init', '8b4d2c82-4277-4380-9212-e4e7f03c1ea4');
  window.Beacon('config', {display: {style: 'manual'}})
</script>
  <script charset="utf-8" src="/js/18.21f55485bf0cacc98665.js"></script><style hs-beacon="active" data-styled-version="5.1.1"></style><script charset="utf-8" src="/js/0.d0e59d3fa10fd6f6b4bd.js"></script><script charset="utf-8" src="/js/1.9d9c2080dba52ad2198a.js"></script><script charset="utf-8" src="/js/7.7f1a83df91855b1a3200.js"></script></head>
  <body>



<div id="vue-container"><div id="top" class="app-container size-sm tablet:size-md icon-discovery"><!----> <!----> <!----> <header id="header" class="app-header padding-y-md" style="background-color: var(--white); --button-background:transparent; --button-hover-background:transparent; --button-margin-bottom:0; --button-padding:var(--spacing-2xs) var(--spacing-md); --button-color:var(--fa-md-gravy); --button-active-color:var(--fa-dk-blue); --button-hover-color:var(--fa-dk-blue); --link-color:var(--fa-md-gravy); --link-hover-color:var(--fa-dk-blue);"><div class="container"><div class="wrap-app-header position-relative display-flex flex-items-center flex-content-between size-lg tablet:size-md"><div class="app-header-menu-toggle display-flex flex-items-center tablet:display-none"><button aria-label="Open Main Menu" data-balloon-pos="right" class="flat swap-icons"><svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="bars" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="initial-icon svg-inline--fa fa-bars fa-xl"><path fill="currentColor" d="M0 88C0 74.75 10.75 64 24 64H424C437.3 64 448 74.75 448 88C448 101.3 437.3 112 424 112H24C10.75 112 0 101.3 0 88zM0 248C0 234.7 10.75 224 24 224H424C437.3 224 448 234.7 448 248C448 261.3 437.3 272 424 272H24C10.75 272 0 261.3 0 248zM424 432H24C10.75 432 0 421.3 0 408C0 394.7 10.75 384 24 384H424C437.3 384 448 394.7 448 408C448 421.3 437.3 432 424 432z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="burger" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="hover-icon svg-inline--fa fa-burger fa-xl"><path fill="currentColor" d="M50.39 220.8C45.93 218.6 42.03 215.5 38.97 211.6C35.91 207.7 33.79 203.2 32.75 198.4C31.71 193.5 31.8 188.5 32.99 183.7C54.98 97.02 146.5 32 256 32C365.5 32 457 97.02 479 183.7C480.2 188.5 480.3 193.5 479.2 198.4C478.2 203.2 476.1 207.7 473 211.6C469.1 215.5 466.1 218.6 461.6 220.8C457.2 222.9 452.3 224 447.3 224H64.67C59.73 224 54.84 222.9 50.39 220.8zM86.4 176H425.6C413 149.5 393 127.3 368 112C368 116.2 366.3 120.3 363.3 123.3C360.3 126.3 356.2 128 352 128C347.8 128 343.7 126.3 340.7 123.3C337.7 120.3 336 116.2 336 112C336 109.3 336.7 106.7 338.1 104.4C339.4 102 341.3 100.1 343.6 98.72C320.4 88.44 295.5 82.41 270.2 80.9C271.3 83.09 271.1 85.52 272 88C272 92.24 270.3 96.31 267.3 99.31C264.3 102.3 260.2 104 256 104C251.8 104 247.7 102.3 244.7 99.31C241.7 96.31 240 92.24 240 88C240 85.52 240.7 83.09 241.8 80.9C216.5 82.41 191.6 88.44 168.4 98.72C170.7 100.1 172.6 102 173.9 104.4C175.3 106.7 175.1 109.3 176 112C176 116.2 174.3 120.3 171.3 123.3C168.3 126.3 164.2 128 160 128C155.8 128 151.7 126.3 148.7 123.3C145.7 120.3 144 116.2 144 112C118.1 127.3 98.95 149.5 86.4 176zM486.6 265.4C492.6 271.4 496 279.5 496 288C496 296.5 492.6 304.6 486.6 310.6C480.6 316.6 472.5 320 464 320H48C39.51 320 31.37 316.6 25.37 310.6C19.37 304.6 16 296.5 16 288C16 279.5 19.37 271.4 25.37 265.4C31.37 259.4 39.51 256 48 256H464C472.5 256 480.6 259.4 486.6 265.4zM475.3 356.7C478.3 359.7 480 363.8 480 368V384C480 409.5 469.9 433.9 451.9 451.9C433.9 469.9 409.5 480 384 480H128C102.5 480 78.12 469.9 60.12 451.9C42.11 433.9 32 409.5 32 384V368C32 363.8 33.69 359.7 36.69 356.7C39.69 353.7 43.76 352 48 352H464C468.2 352 472.3 353.7 475.3 356.7zM411.7 423.2C419.8 417.5 425.9 409.4 429.2 400H82.75C86.06 409.4 92.19 417.5 100.3 423.2C108.4 428.9 118.1 432 128 432H384C393.9 432 403.6 428.9 411.7 423.2z" class=""></path></svg> <span class="sr-only">Open Main Menu</span></button></div> <div class="app-header-links display-flex flex-items-center tablet:flex-items-stretch flex-content-center tablet:flex-content-start"><a href="/" class="wrapper-logo-flag-animated button flat router-link-active"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" shape-rendering="geometricPrecision" text-rendering="geometricPrecision" class="logo-flag-animated" style="height: var(--size-xl);"><g clip-path="url(#logo-flag-animated-mask-clip-path)"><g transform="translate(11.3,7)" class="logo-flag-animated-flag-wave-to"><path d="M11.25,2.5c.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.01824C18.8946,2.26203,19.5368,2.5,20.5,2.5c.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.01824c.658.24379,1.3002.48176,2.2634.48176.9875,0,1.6438-.25,2.3-.5.6563-.25,1.3125-.5,2.3001-.5.9873,0,1.6624.25,2.3374.5l.0491.01824C37.3946,2.26203,38.0368,2.5,39,2.5v2c-.9666,0-1.616-.23962-2.2791-.48438l-.049-.01794C35.9958,3.74841,35.322,3.5,34.3501,3.5c-.9876,0-1.6438.25-2.3001.5-.6562.25-1.3125.5-2.3.5-.9666,0-1.6161-.23963-2.2791-.48422-.0163-.00603-.0327-.01207-.049-.0181C26.7458,3.74841,26.072,3.5,25.1001,3.5c-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.616-.23962-2.2791-.48438l-.049-.01794C17.4958,3.74841,16.822,3.5,15.8501,3.5c-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.61597-.23962-2.27905-.48438l-.04907-.01794C8.24585,3.74841,7.57202,3.5,6.6001,3.5c-.98755,0-1.6438.25-2.30005.5s-1.3125.5-2.30005.5v-2c.98755,0,1.6438-.25,2.30005-.5s1.3125-.5,2.30005-.5c.9873,0,1.66235.25,2.3374.5l.04834.01794C9.64377,2.26174,10.2868,2.5,11.25,2.5Zm0,8c.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.0182c.658.2438,1.3002.4818,2.2634.4818.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.0182c.658.2438,1.3002.4818,2.2634.4818.9875,0,1.6438-.25,2.3-.5.6563-.25,1.3125-.5,2.3001-.5.9873,0,1.6624.25,2.3374.5l.0491.0182c.658.2438,1.3002.4818,2.2634.4818v2c-.9666,0-1.616-.2396-2.2791-.4844l-.049-.0179c-.6761-.2493-1.3499-.4977-2.3218-.4977-.9876,0-1.6438.25-2.3001.5-.6562.25-1.3125.5-2.3.5-.9666,0-1.6161-.2396-2.2791-.4842-.0163-.0061-.0327-.0121-.049-.0181-.6761-.2493-1.3499-.4977-2.3218-.4977-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.616-.2396-2.2791-.4844l-.049-.0179c-.6761-.2493-1.3499-.4977-2.3218-.4977-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.61597-.2396-2.27905-.4844l-.04907-.0179C8.24585,11.7484,7.57202,11.5,6.6001,11.5c-.98755,0-1.6438.25-2.30005.5s-1.3125.5-2.30005.5v-2c.98755,0,1.6438-.25,2.30005-.5s1.3125-.5,2.30005-.5c.9873,0,1.66235.25,2.3374.5l.04834.0179c.65793.2438,1.30096.4821,2.26416.4821Z" transform="translate(-20.5,-7)" fill="var(--fa-brand-blue)"></path></g> <rect width="2" height="8" rx="0" ry="0" transform="translate(13 3)" fill="var(--fa-brand-blue)"></rect> <path d="M0,16h16v-16L0,0v16ZM2,1q.075195.556641,0,7h3v7h10v-14L2,1Z" clip-rule="evenodd" fill="var(--white)" fill-rule="evenodd"></path> <clipPath id="logo-flag-animated-mask-clip-path"><rect width="16" height="16" rx="0" ry="0" fill="var(--white)"></rect></clipPath></g> <path d="M2,1c-.55228,0-1,.44772-1,1v12c0,.5523.44772,1,1,1s1-.4477,1-1L3,2c0-.55228-.44772-1-1-1Z" fill="var(--fa-navy)"></path></svg> <span class="sr-only">Font Awesome</span></a> <div class="display-none tablet:display-block tablet:margin-left-4xl"><div class="wrap-links display-flex flex-items-center"><a href="/start" class="no-underline padding-y-xs padding-x-md margin-right-md">Start</a> <a href="/search?s=solid%2Cbrands" class="no-underline padding-y-xs padding-x-md margin-right-md active" title="Search icons"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="magnifying-glass" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-magnifying-glass"><path fill="currentColor" d="M500.3 443.7l-119.7-119.7c27.22-40.41 40.65-90.9 33.46-144.7C401.8 87.79 326.8 13.32 235.2 1.723C99.01-15.51-15.51 99.01 1.724 235.2c11.6 91.64 86.08 166.7 177.6 178.9c53.8 7.189 104.3-6.236 144.7-33.46l119.7 119.7c15.62 15.62 40.95 15.62 56.57 0C515.9 484.7 515.9 459.3 500.3 443.7zM79.1 208c0-70.58 57.42-128 128-128s128 57.42 128 128c0 70.58-57.42 128-128 128S79.1 278.6 79.1 208z" class=""></path></svg> <span class="sr-only">Search Icons</span></a> <a href="/icons" class="no-underline padding-y-xs padding-x-md margin-right-md">Icons</a> <a href="/docs" class="no-underline padding-y-xs padding-x-md margin-right-md">Docs</a> <a href="/plans" class="no-underline padding-y-xs padding-x-md margin-right-md">Plans</a> <a href="/support" class="no-underline padding-y-xs padding-x-md margin-right-md">Support</a> <a href="https://blog.fontawesome.com" target="_blank" class="no-underline padding-y-xs padding-x-md margin-right-md">Blog</a></div></div></div> <div class="app-header-account display-flex flex-items-center tablet:margin-left-2xl"><a href="/sessions/sign-in?next=%2Fsearch%3Fp%3D2%26c%3Dmedia-playback" class="button flat swap-icons" aria-label="Sign In" data-balloon-pos="left"><svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="right-to-bracket" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="initial-icon svg-inline--fa fa-right-to-bracket fa-xl"><path fill="currentColor" d="M512 128v256c0 53.02-42.98 96-96 96h-72C330.7 480 320 469.3 320 456c0-13.26 10.75-24 24-24H416c26.4 0 48-21.6 48-48V128c0-26.4-21.6-48-48-48h-72C330.7 80 320 69.25 320 56C320 42.74 330.7 32 344 32H416C469 32 512 74.98 512 128zM367.9 273.9L215.5 407.6C209.3 413.1 201.3 416 193.3 416c-4.688 0-9.406-.9687-13.84-2.969C167.6 407.7 160 396.1 160 383.3V328H40C17.94 328 0 310.1 0 288V224c0-22.06 17.94-40 40-40H160V128.7c0-12.75 7.625-24.41 19.41-29.72C191.5 93.56 205.7 95.69 215.5 104.4l152.4 133.6C373.1 242.6 376 249.1 376 256S373.1 269.4 367.9 273.9zM315.8 256L208 161.1V232h-160v48h160v70.03L315.8 256z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="person-to-portal" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="hover-icon svg-inline--fa fa-person-to-portal fa-xl"><path fill="currentColor" d="M416 0c-38.75 0-83.88 34.88-93.88 188.9c-9.25-30.25-34-53.25-64.88-60.25L179 110.9C153.4 105 126.5 111 105.8 127.1L57.38 164.5c-10.5 8-12.38 23.12-4.25 33.62S76.25 210.5 86.75 202.4L135.1 165c9.5-7.25 21.75-9.875 33.38-7.375L183.4 161L148 248.4c-10.38 25.5-.625 54.75 23 68.87l83.75 50.63c2.375 1.625 3.75 4.125 3.875 6.875c0 .75-.125 1.5-.25 2.25L225 481.4C223.3 487.5 224 494.1 227.1 499.6c3.125 5.625 8.25 9.75 14.38 11.5C243.8 511.6 245.9 512 248.1 512c10.75 0 20.12-7.125 23.12-17.38l33.25-104.5c7-24.37-3.25-50.25-24.88-63.37L227.8 295.5l42-104.7c2.75 3.5 5 7.625 6.375 12l14 46c7.125 23.62 28.63 39 53.38 39.12L391.9 288C405.1 288 416 277.2 416 264.1c0-13.25-10.62-23.5-24-23.62h-23.88C369.1 132.9 390.8 48 416 48c26.5 0 48 93.12 48 208s-21.5 208-48 208c-21.38 0-39.38-60.5-45.63-144h-48.5C328 418.1 350.9 512 416 512c43.75 0 96-44.38 96-256C512 139.2 495.4 0 416 0zM272 96c26.5 0 48-21.5 48-47.1S298.5 0 272 0c-26.5 0-48.01 21.5-48.01 48S245.5 96 272 96zM126.1 316.9L106.2 363.1C105 366.1 102.1 368 99 368H24C10.75 368 0 378.8 0 392S10.75 416 24 416h75c22.38 0 42.62-13.38 51.5-33.1L164 350.5l-9.5-5.875C143 337.6 133.4 328.1 126.1 316.9z" class=""></path></svg> <span class="sr-only">Sign In</span></a></div></div> <!----></div></header>  <main id="searchContent" role="main" class="app-content"><div class="ais-InstantSearch"><!----> <section id="icons-header" class="padding-top-2xl" style="background: var(--white);"><div class="container"><div class="row roomy padding-bottom-3xl"><div class="column-12 laptop:column-10 laptop:offset-1 desktop:column-6 desktop:offset-3"><h1 class="sr-only">Font Awesome Icons</h1> <div class="ais-SearchBox" id="aisSearchBox"><div id="icons-search" class="form icons-search position-relative"><div class="with-icon-before" style="--input-with-icon-color:var(--fa-navy); --input-border-color:var(--fa-navy);"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="magnifying-glass" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-magnifying-glass" style="line-height: 1;"><path fill="currentColor" d="M500.3 443.7l-119.7-119.7c27.22-40.41 40.65-90.9 33.46-144.7C401.8 87.79 326.8 13.32 235.2 1.723C99.01-15.51-15.51 99.01 1.724 235.2c11.6 91.64 86.08 166.7 177.6 178.9c53.8 7.189 104.3-6.236 144.7-33.46l119.7 119.7c15.62 15.62 40.95 15.62 56.57 0C515.9 484.7 515.9 459.3 500.3 443.7zM79.1 208c0-70.58 57.42-128 128-128s128 57.42 128 128c0 70.58-57.42 128-128 128S79.1 278.6 79.1 208z" class=""></path></svg> <label for="icons-search" class="sr-only">Search the v6 Icons</label> <input id="icons-search" name="icons-search" type="text" placeholder="Search 16,083 icons..." class="rounded"></div> <!----> <!----> <div class="icons-version-selector"><select id="icons-version" name="icons-version" class="version-select margin-0 padding-left-md padding-right-2xl" style="--input-background:transparent; --input-border-color:transparent; --input-box-shadow:none; --input-padding-horizontal:var(--spacing-xl); background-position: 90% center;"><option>
        6.1.1
      </option><option>
        5.15.4
      </option><option>
        All Versions
      </option></select></div></div></div> <div class="display-flex flex-column flex-items-center tablet:flex-row tablet:flex-content-between padding-horizontal-2xl padding-top-2xs"><p class="display-none tablet:display-block size-sm tablet:margin-bottom-0"><span class="sr-only">Need a starting search? </span>Try

    <!----> <button class="link flat padding-6xs margin-0 display-inline" style="font-weight: var(--font-weight-normal); vertical-align: baseline;">
        hot
      </button><span>, </span><!----> <button class="link flat padding-6xs margin-0 display-inline" style="font-weight: var(--font-weight-normal); vertical-align: baseline;">
        photo
      </button><span>, </span><!----> <button class="link flat padding-6xs margin-0 display-inline" style="font-weight: var(--font-weight-normal); vertical-align: baseline;">
        open
      </button><span>, </span><!----> <button class="link flat padding-6xs margin-0 display-inline" style="font-weight: var(--font-weight-normal); vertical-align: baseline;">
        spotify
      </button><span>, </span><!----> <button class="link flat padding-6xs margin-0 display-inline" style="font-weight: var(--font-weight-normal); vertical-align: baseline;">
        bluetooth
      </button><span>, </span><span>or </span> <button class="link flat padding-6xs margin-0 display-inline" style="font-weight: var(--font-weight-normal); vertical-align: baseline;">
        geordi la forge
      </button><span>.</span></p> <p class="size-sm muted"><span class="sr-only">This icon search is </span><span class="text-capitalize">powered</span> by
    <a href="https://algolia.com" target="_blank" class="no-underline secondary margin-left-4xs"><svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="algolia" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="svg-inline--fa fa-algolia fa-lg"><path fill="currentColor" d="M229.3 182.6c-49.3 0-89.2 39.9-89.2 89.2 0 49.3 39.9 89.2 89.2 89.2s89.2-39.9 89.2-89.2c0-49.3-40-89.2-89.2-89.2zm62.7 56.6l-58.9 30.6c-1.8 .9-3.8-.4-3.8-2.3V201c0-1.5 1.3-2.7 2.7-2.6 26.2 1 48.9 15.7 61.1 37.1 .7 1.3 .2 3-1.1 3.7zM389.1 32H58.9C26.4 32 0 58.4 0 90.9V421c0 32.6 26.4 59 58.9 59H389c32.6 0 58.9-26.4 58.9-58.9V90.9C448 58.4 421.6 32 389.1 32zm-202.6 84.7c0-10.8 8.7-19.5 19.5-19.5h45.3c10.8 0 19.5 8.7 19.5 19.5v15.4c0 1.8-1.7 3-3.3 2.5-12.3-3.4-25.1-5.1-38.1-5.1-13.5 0-26.7 1.8-39.4 5.5-1.7 .5-3.4-.8-3.4-2.5v-15.8zm-84.4 37l9.2-9.2c7.6-7.6 19.9-7.6 27.5 0l7.7 7.7c1.1 1.1 1 3-.3 4-6.2 4.5-12.1 9.4-17.6 14.9-5.4 5.4-10.4 11.3-14.8 17.4-1 1.3-2.9 1.5-4 .3l-7.7-7.7c-7.6-7.5-7.6-19.8 0-27.4zm127.2 244.8c-70 0-126.6-56.7-126.6-126.6s56.7-126.6 126.6-126.6c70 0 126.6 56.6 126.6 126.6 0 69.8-56.7 126.6-126.6 126.6z" class=""></path></svg> Algolia
    </a></p></div></div></div> <div class="display-flex flex-column laptop:flex-row laptop:flex-items-center flex-content-between padding-bottom-3xl"><div class="icons-facets-group-styles display-flex flex-items-center flex-content-between laptop:flex-content-start ais-RefinementList"><div class="icons-facets-group-styles-classic buttons margin-right-md tablet:margin-right-xl" style="--button-transition-duration:var(--timing-2xfast);"><button aria-label="Solid" data-balloon-pos="up" class="icons-facet icons-facet-button display-flex flex-items-center flex-content-center padding-x-lg tablet:padding-x-xl"><span class="icons-facet-affordance position-relative"><i class="fas fa-icons fa-fw fa-lg"></i></span> <span class="sr-only">Show </span> <span class="display-none laptop:display-inline text-capitalize margin-left-sm">Solid</span> <span class="sr-only"> style icons</span> <span class="icons-facet-count sr-only">
          48 matching results
        </span></button><button aria-label="Regular" data-balloon-pos="up" class="icons-facet icons-facet-button display-flex flex-items-center flex-content-center padding-x-lg tablet:padding-x-xl"><span class="icons-facet-affordance position-relative"><i class="far fa-icons fa-fw fa-lg"></i></span> <span class="sr-only">Show </span> <span class="display-none laptop:display-inline text-capitalize margin-left-sm">Regular</span> <span class="sr-only"> style icons</span> <span class="icons-facet-count sr-only">
          48 matching results
        </span></button><button aria-label="Light" data-balloon-pos="up" class="icons-facet icons-facet-button display-flex flex-items-center flex-content-center padding-x-lg tablet:padding-x-xl"><span class="icons-facet-affordance position-relative"><i class="fal fa-icons fa-fw fa-lg"></i></span> <span class="sr-only">Show </span> <span class="display-none laptop:display-inline text-capitalize margin-left-sm">Light</span> <span class="sr-only"> style icons</span> <span class="icons-facet-count sr-only">
          48 matching results
        </span></button><button aria-label="Thin" data-balloon-pos="up" class="icons-facet icons-facet-button display-flex flex-items-center flex-content-center padding-x-lg tablet:padding-x-xl"><span class="icons-facet-affordance position-relative"><i class="fat fa-icons fa-fw fa-lg"></i></span> <span class="sr-only">Show </span> <span class="display-none laptop:display-inline text-capitalize margin-left-sm">Thin</span> <span class="sr-only"> style icons</span> <span class="icons-facet-count sr-only">
          48 matching results
        </span></button><button aria-label="Duotone" data-balloon-pos="up" class="icons-facet icons-facet-button display-flex flex-items-center flex-content-center padding-x-lg tablet:padding-x-xl"><span class="icons-facet-affordance position-relative"><i class="fad fa-icons fa-fw fa-lg"></i></span> <span class="sr-only">Show </span> <span class="display-none laptop:display-inline text-capitalize margin-left-sm">Duotone</span> <span class="sr-only"> style icons</span> <span class="icons-facet-count sr-only">
          48 matching results
        </span></button></div><div class="" style="--button-transition-duration:var(--timing-2xfast);"><button aria-label="Brands" data-balloon-pos="up" class="icons-facet icons-facet-button display-flex flex-items-center flex-content-center padding-x-lg tablet:padding-x-xl"><span class="icons-facet-affordance position-relative"><i class="fab fa-font-awesome-flag fa-fw fa-lg"></i></span> <span class="sr-only">Show </span> <span class="display-none laptop:display-inline text-capitalize margin-left-sm">Brands</span> <span class="sr-only"> style icons</span> <span class="icons-facet-count sr-only">
          0 matching results
        </span></button></div></div> <div class="display-flex flex-items-center flex-content-between margin-top-2xl laptop:margin-top-0 laptop:margin-left-lg" style="--button-margin-bottom:0; --button-background:transparent; --button-color:var(--fa-navy); --button-hover-background:transparent; --button-hover-color:var(--fa-dk-blue);"><button aria-label="Show Filters" data-balloon-pos="up" class="flat compact icons-facets-menu-toggle laptop:display-none"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="filter" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-filter fa-fw fa-lg"><path fill="currentColor" d="M3.853 54.87C10.47 40.9 24.54 32 40 32H472C487.5 32 501.5 40.9 508.1 54.87C514.8 68.84 512.7 85.37 502.1 97.33L320 320.9V448C320 460.1 313.2 471.2 302.3 476.6C291.5 482 278.5 480.9 268.8 473.6L204.8 425.6C196.7 419.6 192 410.1 192 400V320.9L9.042 97.33C-.745 85.37-2.765 68.84 3.854 54.87L3.853 54.87z" class=""></path></svg>
                Show Filters
              </button> <div class="margin-left-lg" style="--button-margin-bottom:0; --button-background:transparent; --button-color:var(--fa-navy); --button-hover-background:transparent; --button-hover-color:var(--fa-dk-blue); --button-transition-duration:var(--timing-2xfast);"><button aria-label="Roomy" data-balloon-pos="up" class="flat compact"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="grid-2" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="svg-inline--fa fa-grid-2 fa-lg"><path fill="currentColor" d="M192 176C192 202.5 170.5 224 144 224H48C21.49 224 0 202.5 0 176V80C0 53.49 21.49 32 48 32H144C170.5 32 192 53.49 192 80V176zM192 432C192 458.5 170.5 480 144 480H48C21.49 480 0 458.5 0 432V336C0 309.5 21.49 288 48 288H144C170.5 288 192 309.5 192 336V432zM256 80C256 53.49 277.5 32 304 32H400C426.5 32 448 53.49 448 80V176C448 202.5 426.5 224 400 224H304C277.5 224 256 202.5 256 176V80zM448 432C448 458.5 426.5 480 400 480H304C277.5 480 256 458.5 256 432V336C256 309.5 277.5 288 304 288H400C426.5 288 448 309.5 448 336V432z" class=""></path></svg> <span class="sr-only">Show icons in roomy display</span></button> <button aria-label="Compact" data-balloon-pos="up" class="link flat compact"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="grid" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="svg-inline--fa fa-grid fa-lg"><path fill="currentColor" d="M0 72C0 49.91 17.91 32 40 32H88C110.1 32 128 49.91 128 72V120C128 142.1 110.1 160 88 160H40C17.91 160 0 142.1 0 120V72zM0 232C0 209.9 17.91 192 40 192H88C110.1 192 128 209.9 128 232V280C128 302.1 110.1 320 88 320H40C17.91 320 0 302.1 0 280V232zM128 440C128 462.1 110.1 480 88 480H40C17.91 480 0 462.1 0 440V392C0 369.9 17.91 352 40 352H88C110.1 352 128 369.9 128 392V440zM160 72C160 49.91 177.9 32 200 32H248C270.1 32 288 49.91 288 72V120C288 142.1 270.1 160 248 160H200C177.9 160 160 142.1 160 120V72zM288 280C288 302.1 270.1 320 248 320H200C177.9 320 160 302.1 160 280V232C160 209.9 177.9 192 200 192H248C270.1 192 288 209.9 288 232V280zM160 392C160 369.9 177.9 352 200 352H248C270.1 352 288 369.9 288 392V440C288 462.1 270.1 480 248 480H200C177.9 480 160 462.1 160 440V392zM448 120C448 142.1 430.1 160 408 160H360C337.9 160 320 142.1 320 120V72C320 49.91 337.9 32 360 32H408C430.1 32 448 49.91 448 72V120zM320 232C320 209.9 337.9 192 360 192H408C430.1 192 448 209.9 448 232V280C448 302.1 430.1 320 408 320H360C337.9 320 320 302.1 320 280V232zM448 440C448 462.1 430.1 480 408 480H360C337.9 480 320 462.1 320 440V392C320 369.9 337.9 352 360 352H408C430.1 352 448 369.9 448 392V440z" class=""></path></svg> <span class="sr-only">Show icons in compact display</span></button> <button aria-label="Cheatsheet" data-balloon-pos="up" class="flat compact"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="list-ul" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-list-ul fa-lg"><path fill="currentColor" d="M16 96C16 69.49 37.49 48 64 48C90.51 48 112 69.49 112 96C112 122.5 90.51 144 64 144C37.49 144 16 122.5 16 96zM480 64C497.7 64 512 78.33 512 96C512 113.7 497.7 128 480 128H192C174.3 128 160 113.7 160 96C160 78.33 174.3 64 192 64H480zM480 224C497.7 224 512 238.3 512 256C512 273.7 497.7 288 480 288H192C174.3 288 160 273.7 160 256C160 238.3 174.3 224 192 224H480zM480 384C497.7 384 512 398.3 512 416C512 433.7 497.7 448 480 448H192C174.3 448 160 433.7 160 416C160 398.3 174.3 384 192 384H480zM16 416C16 389.5 37.49 368 64 368C90.51 368 112 389.5 112 416C112 442.5 90.51 464 64 464C37.49 464 16 442.5 16 416zM112 256C112 282.5 90.51 304 64 304C37.49 304 16 282.5 16 256C16 229.5 37.49 208 64 208C90.51 208 112 229.5 112 256z" class=""></path></svg> <span class="sr-only">Show icons in cheatsheet display</span></button></div></div></div></div></section> <div class="margin-top-2xl margin-bottom-6xl"><div class="container"><!----> <div class="row roomy"><div class="column-12 laptop:column-3 desktop:column-2 wrap-icons-facets-menu is-hidden"><form id="icons-facets-menu" class="form icons-facets-menu display-flex flex-column"><ul class="icons-facets-group-editorial" style="--list-style-type:none; --padding-left:0; --margin-top:0; --margin-bottom:var(--spacing-xl);"><div class="ais-ToggleRefinement"><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-filter-is_free" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-filter-is_free" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="bolt" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512" class="unchecked-icon svg-inline--fa fa-bolt fa-fw"><path fill="currentColor" d="M240.5 224H352C365.3 224 377.3 232.3 381.1 244.7C386.6 257.2 383.1 271.3 373.1 280.1L117.1 504.1C105.8 513.9 89.27 514.7 77.19 505.9C65.1 497.1 60.7 481.1 66.59 467.4L143.5 288H31.1C18.67 288 6.733 279.7 2.044 267.3C-2.645 254.8 .8944 240.7 10.93 231.9L266.9 7.918C278.2-1.92 294.7-2.669 306.8 6.114C318.9 14.9 323.3 30.87 317.4 44.61L240.5 224z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Free</span> <span class="sr-only"> icons</span></span> <span class="icons-facet-count margin-left-sm"><span>41 </span> <span class="sr-only"> matching results</span></span></label></li></div><div class="ais-ToggleRefinement"><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-filter-is_new_in_v6" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-filter-is_new_in_v6" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="sparkles" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="unchecked-icon svg-inline--fa fa-sparkles fa-fw"><path fill="currentColor" d="M327.5 85.19L384 64L405.2 7.491C406.9 2.985 411.2 0 416 0C420.8 0 425.1 2.985 426.8 7.491L448 64L504.5 85.19C509 86.88 512 91.19 512 96C512 100.8 509 105.1 504.5 106.8L448 128L426.8 184.5C425.1 189 420.8 192 416 192C411.2 192 406.9 189 405.2 184.5L384 128L327.5 106.8C322.1 105.1 320 100.8 320 96C320 91.19 322.1 86.88 327.5 85.19V85.19zM257.8 187.3L371.8 240C377.5 242.6 381.1 248.3 381.1 254.6C381.1 260.8 377.5 266.5 371.8 269.1L257.8 321.8L205.1 435.8C202.5 441.5 196.8 445.1 190.6 445.1C184.3 445.1 178.6 441.5 176 435.8L123.3 321.8L9.292 269.1C3.627 266.5 0 260.8 0 254.6C0 248.3 3.627 242.6 9.292 240L123.3 187.3L176 73.29C178.6 67.63 184.3 64 190.6 64C196.8 64 202.5 67.63 205.1 73.29L257.8 187.3zM405.2 327.5C406.9 322.1 411.2 320 416 320C420.8 320 425.1 322.1 426.8 327.5L448 384L504.5 405.2C509 406.9 512 411.2 512 416C512 420.8 509 425.1 504.5 426.8L448 448L426.8 504.5C425.1 509 420.8 512 416 512C411.2 512 406.9 509 405.2 504.5L384 448L327.5 426.8C322.1 425.1 320 420.8 320 416C320 411.2 322.1 406.9 327.5 405.2L384 384L405.2 327.5z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">New in v6</span> <span class="sr-only"> icons</span></span> <span class="icons-facet-count margin-left-sm"><span>56 </span> <span class="sr-only"> matching results</span></span></label></li></div><div class="ais-ToggleRefinement"><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-filter-is_sponsored" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-filter-is_sponsored" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="circle-dollar" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="unchecked-icon svg-inline--fa fa-circle-dollar fa-fw"><path fill="currentColor" d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM276.8 133.6C276.8 121.7 267.1 112 255.2 112C243.3 112 233.6 121.7 233.6 133.6V149.7C227.5 151.1 221.6 152.1 216.1 155.4C198.1 163.1 181.1 177.6 177.1 199.4C172.9 222.5 182.9 243.3 202.4 255.8C216.6 264.9 235.4 270.2 250.5 274.5C252.7 275.1 254.1 275.7 257.3 276.4C269.2 279.6 281.7 282.1 291.8 289.9C303.5 297.9 297.9 312.1 285.9 316.1C277 320.6 263.4 322.1 246.4 319.6C234.8 317.8 223.2 313.8 211.9 310C209.4 309.2 206.9 308.3 204.4 307.5C193.1 303.7 180.8 309.9 177.1 321.2C173.3 332.5 179.5 344.8 190.8 348.5C204.9 353.1 219.1 357.8 233.6 361L233.6 378.4C233.6 390.3 243.3 400 255.2 400C267.1 400 276.8 390.3 276.8 378.4L276.8 363.2C285.7 362.2 294.2 360.2 302 357C320.8 349.5 336.3 334.8 340.5 312.4C344.7 289.2 335.6 267.5 316.2 254.2C301.3 244 281.4 238.4 265.8 233.9C263.5 233.3 261.1 232.6 258.8 231.1C247.4 228.9 235.5 225.7 225.6 219.4C213.1 211.4 222.2 199.8 233.2 195C245.6 189.7 260.9 188.5 274.1 191C281.2 192.4 288.2 194.4 295.2 196.3C296.7 196.8 298.3 197.2 299.8 197.6C311.3 200.8 323.2 194.1 326.4 182.6C329.6 171.1 322.8 159.2 311.3 155.1C309.5 155.5 307.6 154.9 305.8 154.4C296.2 151.7 286.5 149 276.8 147.8L276.8 133.6z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Sponsored</span> <span class="sr-only"> icons</span></span> <span class="icons-facet-count margin-left-sm"><span>5 </span> <span class="sr-only"> matching results</span></span></label></li></div><div class="ais-ToggleRefinement ais-ToggleRefinement--noRefinement"><li style="display: none;"></li></div></ul> <div class="wrap-ad wrap-carbon-ad card with-close" style="--card-margin-bottom:var(--spacing-3xl);"><div class="carbon-ad text-center size-sm position-relative"><script async="" type="text/javascript" src="//cdn.carbonads.com/carbon.js?zoneid=1673&amp;serve=C6AILKT&amp;placement=fontawesome" id="_carbonads_js"></script></div> <span class="sr-only">Advertisement</span> <button type="button" aria-label="Remove Ad" data-balloon-pos="left" class="close" style="--with-close-close-color:var(--fa-gravy);"><svg aria-hidden="true" focusable="false" data-prefix="fad" data-icon="circle-xmark" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-circle-xmark"><g class="fa-duotone-group"><path fill="currentColor" d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM175 208.1L222.1 255.1L175 303C165.7 312.4 165.7 327.6 175 336.1C184.4 346.3 199.6 346.3 208.1 336.1L255.1 289.9L303 336.1C312.4 346.3 327.6 346.3 336.1 336.1C346.3 327.6 346.3 312.4 336.1 303L289.9 255.1L336.1 208.1C346.3 199.6 346.3 184.4 336.1 175C327.6 165.7 312.4 165.7 303 175L255.1 222.1L208.1 175C199.6 165.7 184.4 165.7 175 175C165.7 184.4 165.7 199.6 175 208.1V208.1z" class="fa-secondary"></path><path fill="currentColor" d="M255.1 222.1L303 175C312.4 165.7 327.6 165.7 336.1 175C346.3 184.4 346.3 199.6 336.1 208.1L289.9 255.1L336.1 303C346.3 312.4 346.3 327.6 336.1 336.1C327.6 346.3 312.4 346.3 303 336.1L255.1 289.9L208.1 336.1C199.6 346.3 184.4 346.3 175 336.1C165.7 327.6 165.7 312.4 175 303L222.1 255.1L175 208.1C165.7 199.6 165.7 184.4 175 175C184.4 165.7 199.6 165.7 208.1 175L255.1 222.1z" class="fa-primary"></path></g></svg></button> <div class="modal modal-remove-ads" style="z-index: var(--depth-foreground); display: none;"><article class="remove-ads-modal card roomy with-close modal-content"><div class="section text-left padding-bottom-lg"><h2>Remove ads with a Pro plan!</h2></div> <div class="section text-left"><div class="message display-flex flex-column tablet:flex-row flex-items-center tablet:flex-items-start margin-top-0 margin-bottom-2xl" style="--message-background:var(--fa-yellow); --message-border-color:var(--fa-yellow);"><span class="tag null null null" style="--tag-background:var(--fa-dk-yellow); --tag-color:var(--fa-navy);"><span class="sr-only">Using this requires Font Awesome Pro</span>
  Pro
</span> <p class="margin-bottom-0 margin-top-md tablet:margin-top-0 tablet:margin-left-md">A subscription to a Pro-level plan will remove all third-party advertisments on fontawesome.com.</p></div> <h3 class="h4 margin-top-0">And of course Pro-level plans come with…</h3> <ul class="fa-ul"><li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="icons" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-icons fa-xl"><path fill="currentColor" d="M500.3 7.251C507.7 13.33 512 22.41 512 31.1V175.1C512 202.5 483.3 223.1 447.1 223.1C412.7 223.1 383.1 202.5 383.1 175.1C383.1 149.5 412.7 127.1 447.1 127.1V71.03L351.1 90.23V207.1C351.1 234.5 323.3 255.1 287.1 255.1C252.7 255.1 223.1 234.5 223.1 207.1C223.1 181.5 252.7 159.1 287.1 159.1V63.1C287.1 48.74 298.8 35.61 313.7 32.62L473.7 .6198C483.1-1.261 492.9 1.173 500.3 7.251H500.3zM74.66 303.1L86.5 286.2C92.43 277.3 102.4 271.1 113.1 271.1H174.9C185.6 271.1 195.6 277.3 201.5 286.2L213.3 303.1H239.1C266.5 303.1 287.1 325.5 287.1 351.1V463.1C287.1 490.5 266.5 511.1 239.1 511.1H47.1C21.49 511.1-.0019 490.5-.0019 463.1V351.1C-.0019 325.5 21.49 303.1 47.1 303.1H74.66zM143.1 359.1C117.5 359.1 95.1 381.5 95.1 407.1C95.1 434.5 117.5 455.1 143.1 455.1C170.5 455.1 191.1 434.5 191.1 407.1C191.1 381.5 170.5 359.1 143.1 359.1zM440.3 367.1H496C502.7 367.1 508.6 372.1 510.1 378.4C513.3 384.6 511.6 391.7 506.5 396L378.5 508C372.9 512.1 364.6 513.3 358.6 508.9C352.6 504.6 350.3 496.6 353.3 489.7L391.7 399.1H336C329.3 399.1 323.4 395.9 321 389.6C318.7 383.4 320.4 376.3 325.5 371.1L453.5 259.1C459.1 255 467.4 254.7 473.4 259.1C479.4 263.4 481.6 271.4 478.7 278.3L440.3 367.1zM116.7 219.1L19.85 119.2C-8.112 90.26-6.614 42.31 24.85 15.34C51.82-8.137 93.26-3.642 118.2 21.83L128.2 32.32L137.7 21.83C162.7-3.642 203.6-8.137 231.6 15.34C262.6 42.31 264.1 90.26 236.1 119.2L139.7 219.1C133.2 225.6 122.7 225.6 116.7 219.1H116.7z" class=""></path></svg></span> <span class="margin-left-lg">All 16,083 Icons in Font Awesome</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="swatchbook" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-swatchbook fa-xl"><path fill="currentColor" d="M0 32C0 14.33 14.33 0 32 0H160C177.7 0 192 14.33 192 32V416C192 469 149 512 96 512C42.98 512 0 469 0 416V32zM128 64H64V128H128V64zM64 256H128V192H64V256zM96 440C109.3 440 120 429.3 120 416C120 402.7 109.3 392 96 392C82.75 392 72 402.7 72 416C72 429.3 82.75 440 96 440zM224 416V154L299.4 78.63C311.9 66.13 332.2 66.13 344.7 78.63L435.2 169.1C447.7 181.6 447.7 201.9 435.2 214.4L223.6 425.9C223.9 422.7 224 419.3 224 416V416zM374.8 320H480C497.7 320 512 334.3 512 352V480C512 497.7 497.7 512 480 512H182.8L374.8 320z" class=""></path></svg></span> <span class="margin-left-lg">Solid, Regular, Light, Thin, and Duotone Styles for Each Icon + Brands</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="file-certificate" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-file-certificate fa-xl"><path fill="currentColor" d="M352 128L352 0H176C149.5 0 128 21.49 128 48l.0007 84.94c1.732-.3867 3.804-1.012 6.679-1.793l3.795-1.031c5.186-1.406 10.54-2.121 15.9-2.121c16.23 0 32.1 6.676 43.55 18.32l2.742 2.793c2.17 2.219 3.688 3.867 5.027 5.102c1.691 .5547 3.809 1.066 6.703 1.855l3.85 1.055c21.2 5.762 37.75 22.53 43.14 43.57l1.124 4.252c.791 3.008 1.293 5.18 1.854 6.91c1.24 1.41 2.92 3.004 5.209 5.324l2.412 2.445c15.48 15.48 21.55 38.49 15.88 59.88l-.852 3.175c-.8926 3.328-1.596 5.645-1.984 7.559C283.4 292 284 294.1 284.8 297.1l1.084 4.125c5.711 21.29-.3828 44.15-15.72 59.48l-2.875 2.914c-2.154 2.18-3.754 3.703-4.949 5.047c-.5605 1.746-1.066 3.941-1.861 6.977l-.9824 3.727c-4.748 18.77-18.02 34.06-35.52 41.47L223.1 512H464c26.51 0 48-21.49 48-48V160h-127.1C366.3 160 352 145.7 352 128zM384 0v128h128L384 0zM228.5 371.5c1.375-5.188 2.303-8.914 3.169-11.82c2.594-8.727 4.549-10.09 15.7-21.43c7.5-7.5 10.38-18.5 7.625-28.75c-5.375-20.62-5.5-17.75 0-38.38c2.75-10.38-.125-21.38-7.625-28.88C236.2 230.9 234.3 229.5 231.7 220.8C230.8 217.9 229.9 214.2 228.5 209C227.8 206.4 226.9 204 225.6 201.8C221.8 195.1 215.5 190.1 207.9 188C195.1 184.5 191.3 183.8 186.6 180C184.8 178.5 182.8 176.5 179.1 173.7C178.6 172.3 176.1 170.6 175.1 168.8C172.6 166.2 169.7 164.4 166.6 162.1C165.8 162.6 165.1 162.3 164.3 162C161.2 160.8 157.8 160 154.4 160c-.0176 0 .0176 0 0 0c-.002 0 .002 0 0 0s.002 0 0 0C151.9 160 149.4 160.3 146.9 161c.2441-.0664 .4809-.0391 .725-.0938C147.4 160.1 147.1 160.9 146.9 161C136.7 163.8 132.3 165.1 127.1 165.1c-4.359 0-8.697-1.375-18.82-4.125C106.7 160.3 104.1 160 101.6 160c-.0176 0 .0176 0 0 0C98.18 160 94.86 160.8 91.68 162C90.89 162.3 90.16 162.6 89.4 162.1C86.31 164.4 83.37 166.2 80.89 168.8C65.89 183.9 68.39 182.4 48.14 188C38.02 190.8 30.1 198.8 27.48 209C25.42 216.8 24.38 221.3 23.05 224.6C20.84 230.3 17.9 232.8 8.606 242.3C6.731 244.1 5.188 246.2 3.906 248.5C.0625 255.2-1.038 263.3 1.024 271.1c5.5 20.62 5.375 17.75 0 38.38c-2.062 7.688-.9616 15.8 2.882 22.54c1.281 2.246 2.825 4.34 4.7 6.215c11.06 11.34 13.09 12.77 15.72 21.48c.875 2.902 1.78 6.613 3.155 11.77c2.625 10.38 10.54 18.38 20.67 21.12c7.188 1.938 10.06 2.688 11.68 3.203s2.003 .7969 4.188 1.797V512l63.99-32l63.99 32v-114.4c4.25-2 1.501-1.125 15.87-5c7.592-2.062 13.91-7.078 17.71-13.84C226.9 376.5 227.8 374.1 228.5 371.5zM128 352c-4.422 0-8.741-.4453-12.91-1.297c-8.336-1.703-16.08-5.027-22.89-9.625c-6.807-4.602-12.68-10.47-17.28-17.28C68.04 313.6 64.02 301.3 64.02 288s4.021-25.58 10.92-35.8c4.596-6.809 10.47-12.68 17.28-17.28c6.809-4.598 14.55-7.922 22.89-9.625C119.3 224.4 123.6 224 128 224c35.37 0 63.99 28.62 63.99 64S163.4 352 128 352z" class=""></path></svg></span> <span class="margin-left-lg">A Perpetual License to Use Pro</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="toolbox" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-toolbox fa-xl"><path fill="currentColor" d="M502.6 182.6l-45.25-45.25C451.4 131.4 443.3 128 434.8 128H384V80C384 53.5 362.5 32 336 32h-160C149.5 32 128 53.5 128 80V128H77.25c-8.5 0-16.62 3.375-22.62 9.375L9.375 182.6C3.375 188.6 0 196.8 0 205.3V304h128v-32C128 263.1 135.1 256 144 256h32C184.9 256 192 263.1 192 272v32h128v-32C320 263.1 327.1 256 336 256h32C376.9 256 384 263.1 384 272v32h128V205.3C512 196.8 508.6 188.6 502.6 182.6zM336 128h-160V80h160V128zM384 368c0 8.875-7.125 16-16 16h-32c-8.875 0-16-7.125-16-16v-32H192v32C192 376.9 184.9 384 176 384h-32C135.1 384 128 376.9 128 368v-32H0V448c0 17.62 14.38 32 32 32h448c17.62 0 32-14.38 32-32v-112h-128V368z" class=""></path></svg></span> <span class="margin-left-lg">Services and Tools to Make Easy Work of Using Icons</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="parachute-box" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-parachute-box fa-xl"><path fill="currentColor" d="M272 192V320H304C311 320 317.7 321.5 323.7 324.2L443.8 192H415.5C415.8 186.7 416 181.4 416 176C416 112.1 393.8 54.84 358.9 16.69C450 49.27 493.4 122.6 507.8 173.6C510.5 183.1 502.1 192 493.1 192H487.1L346.8 346.3C350.1 352.8 352 360.2 352 368V464C352 490.5 330.5 512 304 512H207.1C181.5 512 159.1 490.5 159.1 464V368C159.1 360.2 161.9 352.8 165.2 346.3L24.92 192H18.89C9 192 1.483 183.1 4.181 173.6C18.64 122.6 61.97 49.27 153.1 16.69C118.2 54.84 96 112.1 96 176C96 181.4 96.16 186.7 96.47 192H68.17L188.3 324.2C194.3 321.5 200.1 320 207.1 320H239.1V192H128.5C128.2 186.7 127.1 181.4 127.1 176C127.1 125 143.9 80.01 168.2 48.43C192.5 16.89 223.8 0 255.1 0C288.2 0 319.5 16.89 343.8 48.43C368.1 80.01 384 125 384 176C384 181.4 383.8 186.7 383.5 192H272z" class=""></path></svg></span> <span class="margin-left-lg">Fresh Icons, Features, and Software Updates</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="user-headset" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="svg-inline--fa fa-user-headset fa-xl"><path fill="currentColor" d="M48 224C56.88 224 64 216.9 64 208V192c0-88.25 71.75-160 160-160s160 71.75 160 160v16C384 252.1 348.1 288 304 288h-32c0-17.62-14.38-32-32-32h-32c-17.62 0-32 14.38-32 32s14.38 32 32 32h96c61.88-.125 111.9-50.13 112-112V192c0-105.9-86.13-192-192-192S32 86.13 32 192v16C32 216.9 39.13 224 48 224zM208 224h32c22.88 0 43.98 12.2 55.36 31.95L304 256c26.5 0 48-21.5 48-47.1V192c0-70.75-57.25-128-128-128s-128 57.25-128 128c0 40.38 19.12 75.99 48.37 99.49C144.2 290.2 144 289.3 144 288C144 252.6 172.6 224 208 224zM314.7 352H133.3C59.7 352 0 411.7 0 485.3C0 500.1 11.94 512 26.66 512H421.3C436.1 512 448 500.1 448 485.3C448 411.7 388.3 352 314.7 352z" class=""></path></svg></span> <span class="margin-left-lg">Plus Support From Real Humans</span></li></ul></div> <div class="section text-left"><button class="block yellow size-lg"><svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="rocket-launch" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="margin-right-sm svg-inline--fa fa-rocket-launch fa-lg"><path fill="currentColor" d="M408 143.1C408 166.1 390.1 183.1 368 183.1C345.9 183.1 328 166.1 328 143.1C328 121.9 345.9 103.1 368 103.1C390.1 103.1 408 121.9 408 143.1zM384 312.1V394.2C384 419.7 370.6 443.2 348.7 456.2L260.2 508.6C252.8 513 243.6 513.1 236.1 508.9C228.6 504.6 224 496.6 224 488V373.3C224 350.6 215 328.1 199 312.1C183 296.1 161.4 288 138.7 288H24C15.38 288 7.414 283.4 3.146 275.9C-1.123 268.4-1.042 259.2 3.357 251.8L55.83 163.3C68.79 141.4 92.33 127.1 117.8 127.1H199.9C281.7-3.798 408.8-8.546 483.9 5.272C495.6 7.411 504.6 16.45 506.7 28.07C520.5 103.2 515.8 230.3 384 312.1V312.1zM197.9 253.9C210.8 260.2 222.6 268.7 232.1 279C243.3 289.4 251.8 301.2 258.1 314.1C363.9 284.1 414.8 234.5 439.7 188C464.7 141.3 466.1 90.47 461.7 50.33C421.5 45.02 370.7 47.34 323.1 72.33C277.5 97.16 227.9 148.1 197.9 253.9H197.9zM41.98 345.5C76.37 311.1 132.1 311.1 166.5 345.5C200.9 379.9 200.9 435.6 166.5 470C117 519.5 .4765 511.5 .4765 511.5C.4765 511.5-7.516 394.1 41.98 345.5V345.5zM64.58 447.4C64.58 447.4 103.3 450.1 119.8 433.6C131.2 422.2 131.2 403.6 119.8 392.2C108.3 380.8 89.81 380.8 78.38 392.2C61.92 408.7 64.58 447.4 64.58 447.4z" class=""></path></svg>
            Get Font Awesome Pro for only $99/yr
          </button> <p class="size-sm muted text-center margin-top-lg">
          Already have a Pro Plan?
          <a href="/sessions/sign-in?next=%2Fsearch%3Fp%3D2%26c%3Dmedia-playback" class="secondary margin-left-2xs">
            Sign in
            <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="right-to-bracket" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-right-to-bracket"><path fill="currentColor" d="M512 128v256c0 53.02-42.98 96-96 96h-72C330.7 480 320 469.3 320 456c0-13.26 10.75-24 24-24H416c26.4 0 48-21.6 48-48V128c0-26.4-21.6-48-48-48h-72C330.7 80 320 69.25 320 56C320 42.74 330.7 32 344 32H416C469 32 512 74.98 512 128zM367.9 273.9L215.5 407.6C209.3 413.1 201.3 416 193.3 416c-4.688 0-9.406-.9687-13.84-2.969C167.6 407.7 160 396.1 160 383.3V328H40C17.94 328 0 310.1 0 288V224c0-22.06 17.94-40 40-40H160V128.7c0-12.75 7.625-24.41 19.41-29.72C191.5 93.56 205.7 95.69 215.5 104.4l152.4 133.6C373.1 242.6 376 249.1 376 256S373.1 269.4 367.9 273.9zM315.8 256L208 161.1V232h-160v48h160v70.03L315.8 256z" class=""></path></svg></a></p></div> <button type="button" aria-label="Close" data-balloon-pos="left" class="close"><svg aria-hidden="true" focusable="false" data-prefix="fad" data-icon="circle-xmark" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-circle-xmark fa-lg"><g class="fa-duotone-group"><path fill="currentColor" d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM175 208.1L222.1 255.1L175 303C165.7 312.4 165.7 327.6 175 336.1C184.4 346.3 199.6 346.3 208.1 336.1L255.1 289.9L303 336.1C312.4 346.3 327.6 346.3 336.1 336.1C346.3 327.6 346.3 312.4 336.1 303L289.9 255.1L336.1 208.1C346.3 199.6 346.3 184.4 336.1 175C327.6 165.7 312.4 165.7 303 175L255.1 222.1L208.1 175C199.6 165.7 184.4 165.7 175 175C165.7 184.4 165.7 199.6 175 208.1V208.1z" class="fa-secondary"></path><path fill="currentColor" d="M255.1 222.1L303 175C312.4 165.7 327.6 165.7 336.1 175C346.3 184.4 346.3 199.6 336.1 208.1L289.9 255.1L336.1 303C346.3 312.4 346.3 327.6 336.1 336.1C327.6 346.3 312.4 346.3 303 336.1L255.1 289.9L208.1 336.1C199.6 346.3 184.4 346.3 175 336.1C165.7 327.6 165.7 312.4 175 303L222.1 255.1L175 208.1C165.7 199.6 165.7 184.4 175 175C184.4 165.7 199.6 165.7 208.1 175L255.1 222.1z" class="fa-primary"></path></g></svg></button></article></div></div> <div class="ais-RefinementList"><ul class="icons-facets-group-categories" style="--list-style-type:none; --padding-left:0; --margin-top:0; --margin-bottom:0;"><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-accessibility" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-accessibility" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-universal-access fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Accessibility</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-arrows" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-arrows" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-square-up-right fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Arrows</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          75
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-business" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-business" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-briefcase fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Business</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-communication" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-communication" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-message-dots fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Communication</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-connectivity" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-connectivity" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-signal-bars fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Connectivity</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-editing" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-editing" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-pen-line fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Editing</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          20
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-education" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-education" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-graduation-cap fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Education</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-hands" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-hands" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-hand-paper fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Hands</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-maps" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-maps" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-map fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Maps</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          10
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-mathematics" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-mathematics" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-abacus fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Mathematics</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-media-playback" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-media-playback" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-play-pause fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Showing </span> <span class="text-capitalize">Media Playback</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          240
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-music-audio" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-music-audio" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-music fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Music + Audio</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          45
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-photos-images" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-photos-images" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-camera-retro fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Photos + Images</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          15
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-shapes" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-shapes" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-shapes fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Shapes</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li><li class="wrap-icons-facet-input margin-bottom-5xs"><input id="icons-category-spinners" type="checkbox" name="categories" class="input-checkbox-custom"> <label for="icons-category-spinners" class="icons-facet icons-facet-input margin-0 flex-row size-sm flex-items-baseline" style="padding: var(--spacing-2xs) var(--spacing-sm); cursor: pointer; display: flex !important;"><span class="icons-facet-affordance position-relative margin-right-sm"><i class="fas fa-rotate fa-fw unchecked-icon"></i> <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="square" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="hover-icon svg-inline--fa fa-square fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM384 80H64C55.16 80 48 87.16 48 96V416C48 424.8 55.16 432 64 432H384C392.8 432 400 424.8 400 416V96C400 87.16 392.8 80 384 80z" class=""></path></svg> <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="square-check" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="checked-icon svg-inline--fa fa-square-check fa-fw"><path fill="currentColor" d="M384 32C419.3 32 448 60.65 448 96V416C448 451.3 419.3 480 384 480H64C28.65 480 0 451.3 0 416V96C0 60.65 28.65 32 64 32H384zM339.8 211.8C350.7 200.9 350.7 183.1 339.8 172.2C328.9 161.3 311.1 161.3 300.2 172.2L192 280.4L147.8 236.2C136.9 225.3 119.1 225.3 108.2 236.2C97.27 247.1 97.27 264.9 108.2 275.8L172.2 339.8C183.1 350.7 200.9 350.7 211.8 339.8L339.8 211.8z" class=""></path></svg></span> <span style="flex: 1 1 auto;"><span class="sr-only">Show </span> <span class="text-capitalize">Spinners</span><span class="sr-only"> category icons</span></span> <span class="icons-facet-count margin-left-sm">
          5
          <span class="sr-only"> matching results</span></span></label></li></ul></div> <button type="submit" class="sr-only">Update Icon Results</button></form></div> <div class="laptop:column-9 desktop:column-10 column-12"><div class="ais-Stats"><div class="icons-results-summary display-flex flex-row flex-wrap flex-items-center laptop:flex-nowrap margin-bottom-xl"><div class="icons-results-summary-count margin-right-lg"><h1 class="h4 text-nowrap">240 Icons</h1></div> <div class="icons-results-summary-facets margin-top-lg laptop:margin-top-0"><div><ul class="display-inline-flex flex-row flex-wrap" style="--list-style-type:none; --padding-left:0; --margin-top:0; --margin-bottom:calc(var(--spacing-2xs) *-1); --margin-right:calc(var(--spacing-sm) *-1);"><!----> <li><span class="sr-only">categories: </span> <ul class="display-inline-flex flex-row flex-wrap" style="--list-style-type:none; --padding-left:0; --margin-top:0; --margin-bottom:0;"><li class="margin-bottom-2xs margin-right-sm"><a href="#" class="icon-facet-pill tag secondary rounded" style="--link-decoration-style:none; --link-decoration-line:none; --link-hover-decoration-style:none; --link-hover-decoration-line:none; --tag-padding:var(--spacing-sm) var(--spacing-md);">
            Media Playback
            <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="xmark" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512" class="icon-facet-pill-affordance margin-left-4xs svg-inline--fa fa-xmark fa-fw"><path fill="currentColor" d="M310.6 361.4c12.5 12.5 12.5 32.75 0 45.25C304.4 412.9 296.2 416 288 416s-16.38-3.125-22.62-9.375L160 301.3L54.63 406.6C48.38 412.9 40.19 416 32 416S15.63 412.9 9.375 406.6c-12.5-12.5-12.5-32.75 0-45.25l105.4-105.4L9.375 150.6c-12.5-12.5-12.5-32.75 0-45.25s32.75-12.5 45.25 0L160 210.8l105.4-105.4c12.5-12.5 32.75-12.5 45.25 0s12.5 32.75 0 45.25l-105.4 105.4L310.6 361.4z" class=""></path></svg></a></li></ul></li> <li class="margin-bottom-2xs margin-right-sm"><button class="icon-facet-pill icon-facet-pill-clear tag flat rounded" style="--tag-background:transparent; --button-margin-bottom:0; --button-hover-color:var(--fa-dk-blue); --button-active-color:var(--fa-dk-blue); --button-hover-background:transparent; --button-active-background:transparent; --button-active-border-width:var(--border-width-sm); --link-decoration-style:none; --link-decoration-line:none; --link-hover-decoration-style:none; --link-hover-decoration-line:none; --tag-padding:var(--spacing-sm) var(--spacing-md);">
    Reset<span class="sr-only"> All Filters + Queries</span></button></li></ul></div></div> <div class="icons-results-summary-pages margin-left-lg text-right" style="flex: 1 0 auto;"><p class="margin-bottom-0 muted text-nowrap"><span class="sr-only">Currently on </span> <span class="text-capitalize">page</span> 2 of 2
                    </p></div></div></div> <div><div class="ais-StateResults"><div style="display: none;"><div class="ais-CurrentRefinements"><article class="card roomy"><svg aria-hidden="true" focusable="false" data-prefix="fad" data-icon="face-thinking" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-face-thinking fa-4x" style="--fa-secondary-opacity:1; --fa-secondary-color:var(--fa-yellow);"><g class="fa-duotone-group"><path fill="currentColor" d="M224.6 510.1C228.6 504.4 231.9 498.1 234.4 491.3L255.4 433.6L291.1 420.6C315 411.9 329.4 388.6 327.9 364.5C326.6 340.1 311.4 318.4 288.7 308.8L166.2 257.3C158.1 253.8 148.7 257.6 145.3 265.8C141.8 273.9 145.6 283.3 153.8 286.7L236.2 321.4L144 354.1V352C144 321.1 118.9 296 88 296C57.07 296 32 321.1 32 352V380C11.61 343.3 0 301 0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C245.4 512 234.9 511.4 224.6 510.1L224.6 510.1zM176.4 207.1C194 207.1 208.4 193.7 208.4 175.1C208.4 158.3 194 143.1 176.4 143.1C158.7 143.1 144.4 158.3 144.4 175.1C144.4 193.7 158.7 207.1 176.4 207.1zM336.4 159.1C318.7 159.1 304.4 174.3 304.4 191.1C304.4 209.7 318.7 223.1 336.4 223.1C354 223.1 368.4 209.7 368.4 191.1C368.4 174.3 354 159.1 336.4 159.1zM229.6 140.1C236.3 145.9 246.4 145.1 252.1 138.4C257.9 131.7 257.1 121.6 250.4 115.9L237.2 104.5C206.4 78.14 162.3 73.95 127.1 94.08L120.1 98.11C112.4 102.5 109.7 112.3 114.1 119.9C118.5 127.6 128.3 130.3 135.9 125.9L142.1 121.9C166.5 108.4 195.9 111.2 216.4 128.8L229.6 140.1z" class="fa-secondary"></path><path fill="currentColor" d="M135.9 125.9C128.3 130.3 118.5 127.6 114.1 119.9C109.7 112.3 112.4 102.5 120.1 98.11L127.1 94.08C162.3 73.95 206.4 78.14 237.2 104.5L250.4 115.9C257.1 121.6 257.9 131.7 252.1 138.4C246.4 145.1 236.3 145.9 229.6 140.1L216.4 128.8C195.9 111.2 166.5 108.4 142.1 121.9L135.9 125.9zM144.4 176C144.4 158.3 158.7 144 176.4 144C194 144 208.4 158.3 208.4 176C208.4 193.7 194 208 176.4 208C158.7 208 144.4 193.7 144.4 176zM112 400.6L263.8 345.4C276.3 340.9 290 347.3 294.6 359.8C299.1 372.3 292.7 386 280.2 390.6L230.4 408.7L204.3 480.4C197.4 499.4 179.4 512 159.2 512H112C85.49 512 64 490.5 64 464V352C64 338.7 74.75 328 88 328C101.3 328 112 338.7 112 352L112 400.6zM368.4 192C368.4 209.7 354 224 336.4 224C318.7 224 304.4 209.7 304.4 192C304.4 174.3 318.7 160 336.4 160C354 160 368.4 174.3 368.4 192z" class="fa-primary"></path></g></svg> <h1 class="margin-y-2xs">No Icons Found</h1> <p class="size-lg">
              There are no icons that match your current filters. Try removing some of them to get better results.
            </p> <button>
              Clear All Filters + Start Over
            </button></article> <div class="row roomy align-stretch margin-top-md"><div class="column-12 tablet:column-6 desktop:column-4" style="height: 100%;"><a href="https://fontawesome.com/community/leaderboard/new#request" class="card roomy blue"><div class="header"><i class="fa-solid fa-message-question fa-fw fa-4x"></i></div> <div class="section"><h2 class="h4">Request an Icon</h2> <p>Suggest an icon idea to our Open Source community and vote it up the ranks of <a href="/community/leaderboard/new" class="">our icon leaderboard</a>.</p></div></a></div> <div class="column-12 tablet:column-6 desktop:column-4" style="height: 100%;"><a href="/docs/web/use-kits/upload-icons" class="card roomy teal"><div class="header"><i class="fa-solid fa-cloud-arrow-up fa-fw fa-4x"></i></div> <div class="section"><h2 class="h4">Upload an Icon</h2> <p>Upload your own icon to a kit and easily use it just like it an official Font Awesome icon!</p></div></a></div> <div class="column-12 desktop:column-4" style="height: 100%;"><a href="/#commissions" class="card roomy yellow"><div class="header"><i class="fa-solid fa-badge-dollar fa-fw fa-4x"></i></div> <div class="section"><h2 class="h4">Commission an Icon</h2> <p>Absolutely need an icon? Sponsor it and we’ll design it for you and add it to Font Awesome.</p></div></a></div></div></div></div></div></div> <div id="icons-results" class="compact icon-listing margin-top-lg margin-bottom-4xl"><article id="icon-scrubber-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-scrubber fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">scrubber</span></button> <!----> <!----></article><article id="icon-scrubber-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-scrubber fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">scrubber</span></button> <!----> <!----></article><article id="icon-scrubber-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-scrubber fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">scrubber</span></button> <!----> <!----></article><article id="icon-scrubber-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-scrubber fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">scrubber</span></button> <!----> <!----></article><article id="icon-scrubber-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-scrubber fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">scrubber</span></button> <!----> <!----></article><article id="icon-shuffle-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-shuffle fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">shuffle</span></button> <!----> <!----></article><article id="icon-shuffle-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-shuffle fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">shuffle</span></button> <!----> <!----></article><article id="icon-shuffle-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-shuffle fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">shuffle</span></button> <!----> <!----></article><article id="icon-shuffle-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-shuffle fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">shuffle</span></button> <!----> <!----></article><article id="icon-shuffle-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-shuffle fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">shuffle</span></button> <!----> <!----></article><article id="icon-sliders-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-sliders fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders</span></button> <!----> <!----></article><article id="icon-sliders-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-sliders fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders</span></button> <!----> <!----></article><article id="icon-sliders-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-sliders fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders</span></button> <!----> <!----></article><article id="icon-sliders-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-sliders fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders</span></button> <!----> <!----></article><article id="icon-sliders-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-sliders fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders</span></button> <!----> <!----></article><article id="icon-sliders-up-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-sliders-up fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders-up</span></button> <!----> <!----></article><article id="icon-sliders-up-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-sliders-up fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders-up</span></button> <!----> <!----></article><article id="icon-sliders-up-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-sliders-up fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders-up</span></button> <!----> <!----></article><article id="icon-sliders-up-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-sliders-up fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders-up</span></button> <!----> <!----></article><article id="icon-sliders-up-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-sliders-up fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">sliders-up</span></button> <!----> <!----></article><article id="icon-stop-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-stop fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">stop</span></button> <!----> <!----></article><article id="icon-stop-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-stop fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">stop</span></button> <!----> <!----></article><article id="icon-stop-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-stop fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">stop</span></button> <!----> <!----></article><article id="icon-stop-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-stop fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">stop</span></button> <!----> <!----></article><article id="icon-stop-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-stop fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">stop</span></button> <!----> <!----></article><article id="icon-up-right-and-down-left-from-center-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-up-right-and-down-left-from-center fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">up-right-and-down-left-from-center</span></button> <!----> <!----></article><article id="icon-up-right-and-down-left-from-center-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-up-right-and-down-left-from-center fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">up-right-and-down-left-from-center</span></button> <!----> <!----></article><article id="icon-up-right-and-down-left-from-center-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-up-right-and-down-left-from-center fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">up-right-and-down-left-from-center</span></button> <!----> <!----></article><article id="icon-up-right-and-down-left-from-center-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-up-right-and-down-left-from-center fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">up-right-and-down-left-from-center</span></button> <!----> <!----></article><article id="icon-up-right-and-down-left-from-center-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-up-right-and-down-left-from-center fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">up-right-and-down-left-from-center</span></button> <!----> <!----></article><article id="icon-volume-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-volume fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume</span></button> <!----> <!----></article><article id="icon-volume-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-volume fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume</span></button> <!----> <!----></article><article id="icon-volume-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-volume fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume</span></button> <!----> <!----></article><article id="icon-volume-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-volume fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume</span></button> <!----> <!----></article><article id="icon-volume-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-volume fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume</span></button> <!----> <!----></article><article id="icon-volume-high-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-volume-high fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-high</span></button> <!----> <!----></article><article id="icon-volume-high-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-volume-high fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-high</span></button> <!----> <!----></article><article id="icon-volume-high-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-volume-high fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-high</span></button> <!----> <!----></article><article id="icon-volume-high-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-volume-high fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-high</span></button> <!----> <!----></article><article id="icon-volume-high-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-volume-high fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-high</span></button> <!----> <!----></article><article id="icon-volume-low-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-volume-low fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-low</span></button> <!----> <!----></article><article id="icon-volume-low-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-volume-low fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-low</span></button> <!----> <!----></article><article id="icon-volume-low-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-volume-low fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-low</span></button> <!----> <!----></article><article id="icon-volume-low-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-volume-low fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-low</span></button> <!----> <!----></article><article id="icon-volume-low-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-volume-low fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-low</span></button> <!----> <!----></article><article id="icon-volume-off-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-volume-off fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-off</span></button> <!----> <!----></article><article id="icon-volume-off-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-volume-off fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-off</span></button> <!----> <!----></article><article id="icon-volume-off-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-volume-off fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-off</span></button> <!----> <!----></article><article id="icon-volume-off-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-volume-off fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-off</span></button> <!----> <!----></article><article id="icon-volume-off-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-volume-off fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-off</span></button> <!----> <!----></article><article id="icon-volume-slash-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-volume-slash fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-slash</span></button> <!----> <!----></article><article id="icon-volume-slash-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-volume-slash fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-slash</span></button> <!----> <!----></article><article id="icon-volume-slash-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-volume-slash fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-slash</span></button> <!----> <!----></article><article id="icon-volume-slash-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-volume-slash fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-slash</span></button> <!----> <!----></article><article id="icon-volume-slash-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-volume-slash fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-slash</span></button> <!----> <!----></article><article id="icon-volume-xmark-thin" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fat fa-volume-xmark fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-xmark</span></button> <!----> <!----></article><article id="icon-volume-xmark-solid" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fas fa-volume-xmark fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-xmark</span></button> <!----> <!----></article><article id="icon-volume-xmark-regular" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="far fa-volume-xmark fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-xmark</span></button> <!----> <!----></article><article id="icon-volume-xmark-light" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fal fa-volume-xmark fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-xmark</span></button> <!----> <!----></article><article id="icon-volume-xmark-duotone" class="wrap-icon with-top-tag"><button class="icon subtle" style="--button-padding:var(--spacing-xl) var(--spacing-sm) var(--spacing-md) var(--spacing-sm); --button-background:var(--white); --button-hover-background:var(--fa-yellow); --button-color:var(--fa-md-gravy); --button-hover-color:var(--fa-navy); width: 100%; height: 100%; --button-font-weight:var(--font-weight-normal); --button-margin-bottom:0;"><i class="fad fa-volume-xmark fa-fw" style="color: var(--fa-navy);"></i> <span class="icon-name">volume-xmark</span></button> <!----> <!----></article></div> <div class="modal modal-pro-icon" style="z-index: var(--depth-foreground); display: none;"><article class="icon-detail-modal card roomy with-close modal-content"><div><div class="container margin-top-4xl margin-bottom-6xl"><div class="padding-y-4xl padding-x-2xl text-center"><svg aria-hidden="true" focusable="false" data-prefix="fad" data-icon="compact-disc" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-compact-disc fa-spin fa-2x" style="--fa-secondary-opacity:1; --fa-secondary-color:var(--fa-gravy); --fa-animation-duration:1s; --fa-animation-timing:ease-in-out;"><g class="fa-duotone-group"><path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256c0 141.4 114.6 256 256 256s256-114.6 256-256C512 114.6 397.4 0 256 0zM96.15 241.4C95.4 249.6 88.77 256 80.72 256H79.63C70.55 256 63.23 247.1 64.06 238.7c8.299-92.64 82.43-166.6 175.2-174.6C248.3 63.28 256 70.74 256 80.09c0 8.348-6.215 15.27-14.27 15.99C164.7 102.9 103.1 164.3 96.15 241.4zM256 351.1c-53.02 0-96-43-96-95.1s42.98-96 96-96s96 43 96 96S309 351.1 256 351.1z" class="fa-secondary"></path><path fill="currentColor" d="M256 159.1c-53.02 0-96 43-96 96s42.98 95.1 96 95.1s96-43 96-95.1S309 159.1 256 159.1zM256 288C238.3 288 224 273.8 224 256s14.3-32 32-32s32 14.25 32 32S273.7 288 256 288z" class="fa-primary"></path></g></svg> <h2 class="margin-top-0">Loading Icons</h2> <p class="size-lg">Stay on target. Stay on target.</p> <p class="size-sm muted"><span class="sr-only">This </span><span class="text-capitalize">Taking</span> a crazy long time? <a href="mailto:hello@fontawesome.com?subject=I'm having trouble with Font Awesome 6's icons">Let us know</a>.</p></div></div> <!----></div> <!----></article></div> <div class="ais-Pagination"><div class="display-flex tablet:display-none flex-content-center flex-wrap flex-items-center"><a href="/search?c=media-playback" aria-label="Previous Page" class="button subtle"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="angle-left" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 512" class="margin-right-sm svg-inline--fa fa-angle-left fa-lg"><path fill="currentColor" d="M192 448c-8.188 0-16.38-3.125-22.62-9.375l-160-160c-12.5-12.5-12.5-32.75 0-45.25l160-160c12.5-12.5 32.75-12.5 45.25 0s12.5 32.75 0 45.25L77.25 256l137.4 137.4c12.5 12.5 12.5 32.75 0 45.25C208.4 444.9 200.2 448 192 448z" class=""></path></svg>  Prev
      </a> <span style="padding: var(--button-padding); vertical-align: middle; border-style: solid; border-width: 0px; margin-bottom: var(--button-margin-bottom);">Page 2
      </span> <!----></div> <div class="display-none tablet:display-flex flex-content-center flex-wrap flex-items-center pagination-large-screen" style="width: 100%;"><!----> <!----> <!----> <button class="subtle" style="">1</button><button class="flat subtle" style="background: var(--fa-dk-blue); color: var(--white);">2</button> <!----> <!----> <!----></div></div></div></div></div></div></div></main>  <footer id="footer" class="app-footer padding-y-3xl" style="background-color: var(--fa-navy); color: var(--fa-gravy); --link-color:var(--white); --link-hover-color:var(--fa-blue);"><div class="container"><div class="row roomy align-top margin-bottom-xl"><div class="column-12 margin-bottom-2xl laptop:column-5 laptop:margin-bottom-0"><a href="/" class="wrapper-logo-flag-animated display-inline-block router-link-active"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" shape-rendering="geometricPrecision" text-rendering="geometricPrecision" class="logo-flag-animated" style="height: var(--size-xl);"><g clip-path="url(#logo-flag-animated-mask-clip-path)"><g transform="translate(11.3,7)" class="logo-flag-animated-flag-wave-to"><path d="M11.25,2.5c.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.01824C18.8946,2.26203,19.5368,2.5,20.5,2.5c.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.01824c.658.24379,1.3002.48176,2.2634.48176.9875,0,1.6438-.25,2.3-.5.6563-.25,1.3125-.5,2.3001-.5.9873,0,1.6624.25,2.3374.5l.0491.01824C37.3946,2.26203,38.0368,2.5,39,2.5v2c-.9666,0-1.616-.23962-2.2791-.48438l-.049-.01794C35.9958,3.74841,35.322,3.5,34.3501,3.5c-.9876,0-1.6438.25-2.3001.5-.6562.25-1.3125.5-2.3.5-.9666,0-1.6161-.23963-2.2791-.48422-.0163-.00603-.0327-.01207-.049-.0181C26.7458,3.74841,26.072,3.5,25.1001,3.5c-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.616-.23962-2.2791-.48438l-.049-.01794C17.4958,3.74841,16.822,3.5,15.8501,3.5c-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.61597-.23962-2.27905-.48438l-.04907-.01794C8.24585,3.74841,7.57202,3.5,6.6001,3.5c-.98755,0-1.6438.25-2.30005.5s-1.3125.5-2.30005.5v-2c.98755,0,1.6438-.25,2.30005-.5s1.3125-.5,2.30005-.5c.9873,0,1.66235.25,2.3374.5l.04834.01794C9.64377,2.26174,10.2868,2.5,11.25,2.5Zm0,8c.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.0182c.658.2438,1.3002.4818,2.2634.4818.9875,0,1.6438-.25,2.3-.5s1.3125-.5,2.3001-.5c.9873,0,1.6624.25,2.3374.5l.0491.0182c.658.2438,1.3002.4818,2.2634.4818.9875,0,1.6438-.25,2.3-.5.6563-.25,1.3125-.5,2.3001-.5.9873,0,1.6624.25,2.3374.5l.0491.0182c.658.2438,1.3002.4818,2.2634.4818v2c-.9666,0-1.616-.2396-2.2791-.4844l-.049-.0179c-.6761-.2493-1.3499-.4977-2.3218-.4977-.9876,0-1.6438.25-2.3001.5-.6562.25-1.3125.5-2.3.5-.9666,0-1.6161-.2396-2.2791-.4842-.0163-.0061-.0327-.0121-.049-.0181-.6761-.2493-1.3499-.4977-2.3218-.4977-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.616-.2396-2.2791-.4844l-.049-.0179c-.6761-.2493-1.3499-.4977-2.3218-.4977-.9876,0-1.6438.25-2.3001.5s-1.3125.5-2.3.5c-.9666,0-1.61597-.2396-2.27905-.4844l-.04907-.0179C8.24585,11.7484,7.57202,11.5,6.6001,11.5c-.98755,0-1.6438.25-2.30005.5s-1.3125.5-2.30005.5v-2c.98755,0,1.6438-.25,2.30005-.5s1.3125-.5,2.30005-.5c.9873,0,1.66235.25,2.3374.5l.04834.0179c.65793.2438,1.30096.4821,2.26416.4821Z" transform="translate(-20.5,-7)" fill="var(--fa-brand-blue)"></path></g> <rect width="2" height="8" rx="0" ry="0" transform="translate(13 3)" fill="var(--fa-brand-blue)"></rect> <path d="M0,16h16v-16L0,0v16ZM2,1q.075195.556641,0,7h3v7h10v-14L2,1Z" clip-rule="evenodd" fill="var(--fa-navy)" fill-rule="evenodd"></path> <clipPath id="logo-flag-animated-mask-clip-path"><rect width="16" height="16" rx="0" ry="0" fill="var(--fa-navy)"></rect></clipPath></g> <path d="M2,1c-.55228,0-1,.44772-1,1v12c0,.5523.44772,1,1,1s1-.4477,1-1L3,2c0-.55228-.44772-1-1-1Z" fill="var(--white)"></path></svg> <span class="sr-only">Font Awesome</span></a> <div style="color: var(--white);"><h2 class="h5 margin-top-xl">Go Make Something Awesome</h2> <p class="laptop:size-sm margin-bottom-md line-length-lg laptop:line-length-xl">Font Awesome is the internet's icon library and toolkit used by millions of designers, developers, and content creators.</p> <p class="laptop:size-sm">
            Made with
            <svg aria-labelledby="svg-inline--fa-title-tloTQxF7qytE" data-prefix="fas" data-icon="heart" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="margin-x-4xs app-footer-heart-flutter svg-inline--fa fa-heart"><title id="svg-inline--fa-title-tloTQxF7qytE" class="">love</title><path fill="currentColor" d="M0 190.9V185.1C0 115.2 50.52 55.58 119.4 44.1C164.1 36.51 211.4 51.37 244 84.02L256 96L267.1 84.02C300.6 51.37 347 36.51 392.6 44.1C461.5 55.58 512 115.2 512 185.1V190.9C512 232.4 494.8 272.1 464.4 300.4L283.7 469.1C276.2 476.1 266.3 480 256 480C245.7 480 235.8 476.1 228.3 469.1L47.59 300.4C17.23 272.1 .0003 232.4 .0003 190.9L0 190.9z" class=""></path></svg>
            and
            <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="sword-laser" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="margin-x-4xs svg-inline--fa fa-sword-laser"><path fill="currentColor" d="M9.375 412.1c-12.5 12.5-12.5 32.75 0 45.25l45.25 45.25c12.5 12.5 32.75 12.5 45.25 0l79.12-79.12l-90.5-90.5L9.375 412.1zM105.5 327.3l79.25 79.25l22.62-22.62L128.1 304.6L105.5 327.3zM195.1 236.8c-3.125-3.125-8.125-3.125-11.25 0L173.4 248.1c-3.125 3.125-3.125 8.125 0 11.25l79.25 79.25c3.125 3.125 8.125 3.125 11.25 0l11.38-11.38c3.125-3.125 3.125-8.123 0-11.25L195.1 236.8zM139.5 293.4l79.12 79.12l22.62-22.62L162.1 270.8L139.5 293.4zM504.1 7.02c-9.125-9.125-23.87-9.375-33.25-.625L225.6 232.4l53.1 53.1l225.1-246.1C514.4 30.89 514.1 16.14 504.1 7.02z" class=""></path></svg>
            in
            <a href="https://www.google.com/maps/place/Bentonville, AR" target="_blank" class="no-underline">Bentonville</a>,
            <a href="https://www.google.com/maps/place/Boston, MA" target="_blank" class="no-underline">Boston</a>,
            <a href="https://www.google.com/maps/place/Chicago, IL" target="_blank" class="no-underline">Chicago</a>,
            <a href="https://www.google.com/maps/place/Grand Rapids, MI" target="_blank" class="no-underline">Grand Rapids</a>,
            <a href="https://www.google.com/maps/place/Joplin, MO" target="_blank" class="no-underline">Joplin</a>,
            <a href="https://www.google.com/maps/place/Kansas City, MO" target="_blank" class="no-underline">Kansas City</a>,
            <a href="https://www.google.com/maps/place/Seattle, WA" target="_blank" class="no-underline">Seattle</a>,
            <a href="https://www.google.com/maps/place/Tampa, FL" target="_blank" class="no-underline">Tampa</a>,
            and <a href="https://www.google.com/maps/place/Vergennes, VT" target="_blank" class="no-underline">Vergennes</a>.
          </p></div></div> <div class="column-12 laptop:column-6 laptop:offset-1"><div class="row roomy align-top"><div class="column-4"><h3 class="h5 muted">Project</h3> <a href="/docs/changelog" class="display-block no-underline margin-bottom-2xs">Changelog</a> <a href="https://status.fortawesome.com/" target="_blank" class="display-block no-underline margin-bottom-2xs">Status</a> <a href="/#commissions" class="display-block no-underline margin-bottom-2xs">Commission Icons</a> <a href="/license" class="display-block no-underline margin-bottom-2xs">License</a> <a href="/versions" class="display-block no-underline">All Versions</a></div> <div class="column-4"><h3 class="h5 muted">Help</h3> <a href="/support" class="display-block no-underline margin-bottom-2xs">FAQs</a> <a href="/docs/web/troubleshoot" class="display-block no-underline margin-bottom-2xs">Troubleshooting</a> <button class="flat link no-underline margin-bottom-2xs" style="--button-padding:0; --button-active-background:transparent; vertical-align: baseline;">
              Support
            </button> <a href="mailto:hello@fontawesome.com?subject=Hello, Font Awesome!" class="display-block no-underline">Contact Us</a></div> <div class="column-4"><h3 class="h5 muted">Community</h3> <a href="https://github.com/FortAwesome/Font-Awesome" target="_blank" class="display-block no-underline margin-bottom-2xs">GitHub</a> <a href="https://github.com/FortAwesome/Font-Awesome/issues" target="_blank" class="display-block no-underline margin-bottom-2xs">Issues</a> <a href="/community/leaderboard/new" class="display-block no-underline margin-bottom-2xs">Icon Requests</a> <a href="https://twitter.com/fontawesome/" target="_blank" class="display-block no-underline">Twitter</a></div></div></div></div> <div class="row roomy align-top"><div class="display-flex flex-row flex-items-center flex-content-end laptop:size-sm" style="width: 100%;"><a href="/tos" class="no-underline padding-x-2xs margin-right-md">Terms of Service</a> <a href="/privacy" class="no-underline padding-x-2xs margin-right-md">Privacy Policy</a> <p style="color: var(--fa-gravy);">© Fonticons, Inc.</p></div></div></div> <div class="modal modal-remove-ads" style="z-index: var(--depth-foreground); display: none;"><article class="card roomy with-close modal-content"><div class="section text-left paddinng-bottom-lg"><h2>Get Official Support with a Pro Plan!</h2></div> <div class="section text-left"><div class="message display-flex flex-column tablet:flex-row flex-items-center tablet:flex-items-start margin-top-0 margin-bottom-2xl" style="--message-background:var(--fa-yellow); --message-border-color:var(--fa-yellow);"><span class="tag null null null" style="--tag-background:var(--fa-dk-yellow); --tag-color:var(--fa-navy);"><span class="sr-only">Using this requires Font Awesome Pro</span>
  Pro
</span> <p class="margin-bottom-0 margin-top-md tablet:margin-top-0 tablet:margin-left-md">A subscription to a Pro-level plan will get you access to Font Awesome team members who will help you troubleshoot and use Font Awesome in your projects.</p></div> <h3 class="h4 margin-top-0">And of course Pro-level plans come with…</h3> <ul class="fa-ul"><li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="icons" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-icons fa-xl"><path fill="currentColor" d="M500.3 7.251C507.7 13.33 512 22.41 512 31.1V175.1C512 202.5 483.3 223.1 447.1 223.1C412.7 223.1 383.1 202.5 383.1 175.1C383.1 149.5 412.7 127.1 447.1 127.1V71.03L351.1 90.23V207.1C351.1 234.5 323.3 255.1 287.1 255.1C252.7 255.1 223.1 234.5 223.1 207.1C223.1 181.5 252.7 159.1 287.1 159.1V63.1C287.1 48.74 298.8 35.61 313.7 32.62L473.7 .6198C483.1-1.261 492.9 1.173 500.3 7.251H500.3zM74.66 303.1L86.5 286.2C92.43 277.3 102.4 271.1 113.1 271.1H174.9C185.6 271.1 195.6 277.3 201.5 286.2L213.3 303.1H239.1C266.5 303.1 287.1 325.5 287.1 351.1V463.1C287.1 490.5 266.5 511.1 239.1 511.1H47.1C21.49 511.1-.0019 490.5-.0019 463.1V351.1C-.0019 325.5 21.49 303.1 47.1 303.1H74.66zM143.1 359.1C117.5 359.1 95.1 381.5 95.1 407.1C95.1 434.5 117.5 455.1 143.1 455.1C170.5 455.1 191.1 434.5 191.1 407.1C191.1 381.5 170.5 359.1 143.1 359.1zM440.3 367.1H496C502.7 367.1 508.6 372.1 510.1 378.4C513.3 384.6 511.6 391.7 506.5 396L378.5 508C372.9 512.1 364.6 513.3 358.6 508.9C352.6 504.6 350.3 496.6 353.3 489.7L391.7 399.1H336C329.3 399.1 323.4 395.9 321 389.6C318.7 383.4 320.4 376.3 325.5 371.1L453.5 259.1C459.1 255 467.4 254.7 473.4 259.1C479.4 263.4 481.6 271.4 478.7 278.3L440.3 367.1zM116.7 219.1L19.85 119.2C-8.112 90.26-6.614 42.31 24.85 15.34C51.82-8.137 93.26-3.642 118.2 21.83L128.2 32.32L137.7 21.83C162.7-3.642 203.6-8.137 231.6 15.34C262.6 42.31 264.1 90.26 236.1 119.2L139.7 219.1C133.2 225.6 122.7 225.6 116.7 219.1H116.7z" class=""></path></svg></span> <span class="margin-left-lg">All 16,083 Icons in Font Awesome</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="swatchbook" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-swatchbook fa-xl"><path fill="currentColor" d="M0 32C0 14.33 14.33 0 32 0H160C177.7 0 192 14.33 192 32V416C192 469 149 512 96 512C42.98 512 0 469 0 416V32zM128 64H64V128H128V64zM64 256H128V192H64V256zM96 440C109.3 440 120 429.3 120 416C120 402.7 109.3 392 96 392C82.75 392 72 402.7 72 416C72 429.3 82.75 440 96 440zM224 416V154L299.4 78.63C311.9 66.13 332.2 66.13 344.7 78.63L435.2 169.1C447.7 181.6 447.7 201.9 435.2 214.4L223.6 425.9C223.9 422.7 224 419.3 224 416V416zM374.8 320H480C497.7 320 512 334.3 512 352V480C512 497.7 497.7 512 480 512H182.8L374.8 320z" class=""></path></svg></span> <span class="margin-left-lg">Solid, Regular, Light, Thin, and Duotone Styles for Each Icon + Brands</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="file-certificate" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-file-certificate fa-xl"><path fill="currentColor" d="M352 128L352 0H176C149.5 0 128 21.49 128 48l.0007 84.94c1.732-.3867 3.804-1.012 6.679-1.793l3.795-1.031c5.186-1.406 10.54-2.121 15.9-2.121c16.23 0 32.1 6.676 43.55 18.32l2.742 2.793c2.17 2.219 3.688 3.867 5.027 5.102c1.691 .5547 3.809 1.066 6.703 1.855l3.85 1.055c21.2 5.762 37.75 22.53 43.14 43.57l1.124 4.252c.791 3.008 1.293 5.18 1.854 6.91c1.24 1.41 2.92 3.004 5.209 5.324l2.412 2.445c15.48 15.48 21.55 38.49 15.88 59.88l-.852 3.175c-.8926 3.328-1.596 5.645-1.984 7.559C283.4 292 284 294.1 284.8 297.1l1.084 4.125c5.711 21.29-.3828 44.15-15.72 59.48l-2.875 2.914c-2.154 2.18-3.754 3.703-4.949 5.047c-.5605 1.746-1.066 3.941-1.861 6.977l-.9824 3.727c-4.748 18.77-18.02 34.06-35.52 41.47L223.1 512H464c26.51 0 48-21.49 48-48V160h-127.1C366.3 160 352 145.7 352 128zM384 0v128h128L384 0zM228.5 371.5c1.375-5.188 2.303-8.914 3.169-11.82c2.594-8.727 4.549-10.09 15.7-21.43c7.5-7.5 10.38-18.5 7.625-28.75c-5.375-20.62-5.5-17.75 0-38.38c2.75-10.38-.125-21.38-7.625-28.88C236.2 230.9 234.3 229.5 231.7 220.8C230.8 217.9 229.9 214.2 228.5 209C227.8 206.4 226.9 204 225.6 201.8C221.8 195.1 215.5 190.1 207.9 188C195.1 184.5 191.3 183.8 186.6 180C184.8 178.5 182.8 176.5 179.1 173.7C178.6 172.3 176.1 170.6 175.1 168.8C172.6 166.2 169.7 164.4 166.6 162.1C165.8 162.6 165.1 162.3 164.3 162C161.2 160.8 157.8 160 154.4 160c-.0176 0 .0176 0 0 0c-.002 0 .002 0 0 0s.002 0 0 0C151.9 160 149.4 160.3 146.9 161c.2441-.0664 .4809-.0391 .725-.0938C147.4 160.1 147.1 160.9 146.9 161C136.7 163.8 132.3 165.1 127.1 165.1c-4.359 0-8.697-1.375-18.82-4.125C106.7 160.3 104.1 160 101.6 160c-.0176 0 .0176 0 0 0C98.18 160 94.86 160.8 91.68 162C90.89 162.3 90.16 162.6 89.4 162.1C86.31 164.4 83.37 166.2 80.89 168.8C65.89 183.9 68.39 182.4 48.14 188C38.02 190.8 30.1 198.8 27.48 209C25.42 216.8 24.38 221.3 23.05 224.6C20.84 230.3 17.9 232.8 8.606 242.3C6.731 244.1 5.188 246.2 3.906 248.5C.0625 255.2-1.038 263.3 1.024 271.1c5.5 20.62 5.375 17.75 0 38.38c-2.062 7.688-.9616 15.8 2.882 22.54c1.281 2.246 2.825 4.34 4.7 6.215c11.06 11.34 13.09 12.77 15.72 21.48c.875 2.902 1.78 6.613 3.155 11.77c2.625 10.38 10.54 18.38 20.67 21.12c7.188 1.938 10.06 2.688 11.68 3.203s2.003 .7969 4.188 1.797V512l63.99-32l63.99 32v-114.4c4.25-2 1.501-1.125 15.87-5c7.592-2.062 13.91-7.078 17.71-13.84C226.9 376.5 227.8 374.1 228.5 371.5zM128 352c-4.422 0-8.741-.4453-12.91-1.297c-8.336-1.703-16.08-5.027-22.89-9.625c-6.807-4.602-12.68-10.47-17.28-17.28C68.04 313.6 64.02 301.3 64.02 288s4.021-25.58 10.92-35.8c4.596-6.809 10.47-12.68 17.28-17.28c6.809-4.598 14.55-7.922 22.89-9.625C119.3 224.4 123.6 224 128 224c35.37 0 63.99 28.62 63.99 64S163.4 352 128 352z" class=""></path></svg></span> <span class="margin-left-lg">A Perpetual License to Use Pro</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="toolbox" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-toolbox fa-xl"><path fill="currentColor" d="M502.6 182.6l-45.25-45.25C451.4 131.4 443.3 128 434.8 128H384V80C384 53.5 362.5 32 336 32h-160C149.5 32 128 53.5 128 80V128H77.25c-8.5 0-16.62 3.375-22.62 9.375L9.375 182.6C3.375 188.6 0 196.8 0 205.3V304h128v-32C128 263.1 135.1 256 144 256h32C184.9 256 192 263.1 192 272v32h128v-32C320 263.1 327.1 256 336 256h32C376.9 256 384 263.1 384 272v32h128V205.3C512 196.8 508.6 188.6 502.6 182.6zM336 128h-160V80h160V128zM384 368c0 8.875-7.125 16-16 16h-32c-8.875 0-16-7.125-16-16v-32H192v32C192 376.9 184.9 384 176 384h-32C135.1 384 128 376.9 128 368v-32H0V448c0 17.62 14.38 32 32 32h448c17.62 0 32-14.38 32-32v-112h-128V368z" class=""></path></svg></span> <span class="margin-left-lg">Services and Tools to Make Easy Work of Using Icons</span></li> <li class="padding-vertical-sm"><span class="fa-li"><svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="parachute-box" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-parachute-box fa-xl"><path fill="currentColor" d="M272 192V320H304C311 320 317.7 321.5 323.7 324.2L443.8 192H415.5C415.8 186.7 416 181.4 416 176C416 112.1 393.8 54.84 358.9 16.69C450 49.27 493.4 122.6 507.8 173.6C510.5 183.1 502.1 192 493.1 192H487.1L346.8 346.3C350.1 352.8 352 360.2 352 368V464C352 490.5 330.5 512 304 512H207.1C181.5 512 159.1 490.5 159.1 464V368C159.1 360.2 161.9 352.8 165.2 346.3L24.92 192H18.89C9 192 1.483 183.1 4.181 173.6C18.64 122.6 61.97 49.27 153.1 16.69C118.2 54.84 96 112.1 96 176C96 181.4 96.16 186.7 96.47 192H68.17L188.3 324.2C194.3 321.5 200.1 320 207.1 320H239.1V192H128.5C128.2 186.7 127.1 181.4 127.1 176C127.1 125 143.9 80.01 168.2 48.43C192.5 16.89 223.8 0 255.1 0C288.2 0 319.5 16.89 343.8 48.43C368.1 80.01 384 125 384 176C384 181.4 383.8 186.7 383.5 192H272z" class=""></path></svg></span> <span class="margin-left-lg">Fresh Icons, Features, and Software Updates</span></li></ul></div> <div class="section text-left"><button class="block yellow size-lg"><svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="rocket-launch" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="margin-right-sm svg-inline--fa fa-rocket-launch fa-lg"><path fill="currentColor" d="M408 143.1C408 166.1 390.1 183.1 368 183.1C345.9 183.1 328 166.1 328 143.1C328 121.9 345.9 103.1 368 103.1C390.1 103.1 408 121.9 408 143.1zM384 312.1V394.2C384 419.7 370.6 443.2 348.7 456.2L260.2 508.6C252.8 513 243.6 513.1 236.1 508.9C228.6 504.6 224 496.6 224 488V373.3C224 350.6 215 328.1 199 312.1C183 296.1 161.4 288 138.7 288H24C15.38 288 7.414 283.4 3.146 275.9C-1.123 268.4-1.042 259.2 3.357 251.8L55.83 163.3C68.79 141.4 92.33 127.1 117.8 127.1H199.9C281.7-3.798 408.8-8.546 483.9 5.272C495.6 7.411 504.6 16.45 506.7 28.07C520.5 103.2 515.8 230.3 384 312.1V312.1zM197.9 253.9C210.8 260.2 222.6 268.7 232.1 279C243.3 289.4 251.8 301.2 258.1 314.1C363.9 284.1 414.8 234.5 439.7 188C464.7 141.3 466.1 90.47 461.7 50.33C421.5 45.02 370.7 47.34 323.1 72.33C277.5 97.16 227.9 148.1 197.9 253.9H197.9zM41.98 345.5C76.37 311.1 132.1 311.1 166.5 345.5C200.9 379.9 200.9 435.6 166.5 470C117 519.5 .4765 511.5 .4765 511.5C.4765 511.5-7.516 394.1 41.98 345.5V345.5zM64.58 447.4C64.58 447.4 103.3 450.1 119.8 433.6C131.2 422.2 131.2 403.6 119.8 392.2C108.3 380.8 89.81 380.8 78.38 392.2C61.92 408.7 64.58 447.4 64.58 447.4z" class=""></path></svg>Get Font Awesome Pro for only $99/yr</button> <p class="size-sm muted text-center margin-top-lg">
          Already have a Pro Plan?
          <a href="/sessions/sign-in?next=%2Fsearch%3Fp%3D2%26c%3Dmedia-playback" class="secondary margin-left-2xs swap-icons">
            Sign in
            <svg aria-hidden="true" focusable="false" data-prefix="far" data-icon="right-to-bracket" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-right-to-bracket"><path fill="currentColor" d="M512 128v256c0 53.02-42.98 96-96 96h-72C330.7 480 320 469.3 320 456c0-13.26 10.75-24 24-24H416c26.4 0 48-21.6 48-48V128c0-26.4-21.6-48-48-48h-72C330.7 80 320 69.25 320 56C320 42.74 330.7 32 344 32H416C469 32 512 74.98 512 128zM367.9 273.9L215.5 407.6C209.3 413.1 201.3 416 193.3 416c-4.688 0-9.406-.9687-13.84-2.969C167.6 407.7 160 396.1 160 383.3V328H40C17.94 328 0 310.1 0 288V224c0-22.06 17.94-40 40-40H160V128.7c0-12.75 7.625-24.41 19.41-29.72C191.5 93.56 205.7 95.69 215.5 104.4l152.4 133.6C373.1 242.6 376 249.1 376 256S373.1 269.4 367.9 273.9zM315.8 256L208 161.1V232h-160v48h160v70.03L315.8 256z" class=""></path></svg></a></p></div> <button type="button" aria-label="Close" data-balloon-pos="left" class="close"><svg aria-hidden="true" focusable="false" data-prefix="fad" data-icon="circle-xmark" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="svg-inline--fa fa-circle-xmark fa-lg"><g class="fa-duotone-group"><path fill="currentColor" d="M0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256zM175 208.1L222.1 255.1L175 303C165.7 312.4 165.7 327.6 175 336.1C184.4 346.3 199.6 346.3 208.1 336.1L255.1 289.9L303 336.1C312.4 346.3 327.6 346.3 336.1 336.1C346.3 327.6 346.3 312.4 336.1 303L289.9 255.1L336.1 208.1C346.3 199.6 346.3 184.4 336.1 175C327.6 165.7 312.4 165.7 303 175L255.1 222.1L208.1 175C199.6 165.7 184.4 165.7 175 175C165.7 184.4 165.7 199.6 175 208.1V208.1z" class="fa-secondary"></path><path fill="currentColor" d="M255.1 222.1L303 175C312.4 165.7 327.6 165.7 336.1 175C346.3 184.4 346.3 199.6 336.1 208.1L289.9 255.1L336.1 303C346.3 312.4 346.3 327.6 336.1 336.1C327.6 346.3 312.4 346.3 303 336.1L255.1 289.9L208.1 336.1C199.6 346.3 184.4 346.3 175 336.1C165.7 327.6 165.7 312.4 175 303L222.1 255.1L175 208.1C165.7 199.6 165.7 184.4 175 175C184.4 165.7 199.6 165.7 208.1 175L255.1 222.1z" class="fa-primary"></path></g></svg></button></article></div></footer></div><!--function(t,n,i,a){return Ft(e,t,n,i,a,!0)}--></div>
<div id="modal"></div>
<div id="shade"></div>
<script>
  window.__inline_data__ = []
</script>

    <script src="https://use.fortawesome.com/349cfdf6.js"></script><link rel="stylesheet" href="https://fonticons-free-fonticons.netdna-ssl.com/kits/349cfdf6/publications/119263/woff2.css" media="all">

    <script defer="" src="https://m.servedby-buysellads.com/monetization.js"></script>
    <script defer="" src="https://js.stripe.com/v3/"></script>
    <script defer="" src="https://www.google.com/recaptcha/api.js?render=6Lfwy8YZAAAAAOymsOdsZ7xDAG-TFKW_fij1Wnjg"></script>
    <script defer="" src="https://embed.typeform.com/embed.js"></script>
  

<iframe name="__privateStripeMetricsController9050" frameborder="0" allowtransparency="true" scrolling="no" allow="payment *" src="https://js.stripe.com/v3/m-outer-08a68483638f1673180e789f690b2a14.html#url=https%3A%2F%2Ffontawesome.com%2Ficons&amp;title=Icons%20%7C%20Font%20Awesome&amp;referrer=https%3A%2F%2Ffontawesome.com%2Fv5%2Ficons%2Fandroid%3Fs%3Dbrands&amp;muid=d283d1fa-6856-4a25-9940-9f837a43951358231d&amp;sid=58511e87-b008-4e32-9296-cfa9e79c836dd06e3f&amp;version=6&amp;preview=false" aria-hidden="true" tabindex="-1" style="border: none !important; margin: 0px !important; padding: 0px !important; width: 1px !important; min-width: 100% !important; overflow: hidden !important; display: block !important; visibility: hidden !important; position: fixed !important; height: 1px !important; pointer-events: none !important; user-select: none !important;"></iframe><div><div class="grecaptcha-badge" data-style="bottomright" style="width: 256px; height: 60px; display: block; transition: right 0.3s ease 0s; position: fixed; bottom: 14px; right: -186px; box-shadow: gray 0px 0px 5px; border-radius: 2px; overflow: hidden;"><div class="grecaptcha-logo"><iframe title="reCAPTCHA" src="https://www.google.com/recaptcha/api2/anchor?ar=2&amp;k=6Lfwy8YZAAAAAOymsOdsZ7xDAG-TFKW_fij1Wnjg&amp;co=aHR0cHM6Ly9mb250YXdlc29tZS5jb206NDQz&amp;hl=zh-CN&amp;v=nEGwmCAyCoKVn9PSwAGnQWhY&amp;size=invisible&amp;cb=rqoly4l2z4oc" width="256" height="60" role="presentation" name="a-9j2dcfp3yu9g" frameborder="0" scrolling="no" sandbox="allow-forms allow-popups allow-same-origin allow-scripts allow-top-navigation allow-modals allow-popups-to-escape-sandbox allow-storage-access-by-user-activation"></iframe></div><div class="grecaptcha-error"></div><textarea id="g-recaptcha-response-100000" name="g-recaptcha-response" class="g-recaptcha-response" style="width: 250px; height: 40px; border: 1px solid rgb(193, 193, 193); margin: 10px 25px; padding: 0px; resize: none; display: none;"></textarea></div><iframe style="display: none;"></iframe></div><div id="beacon-container"><div class="hsds-beacon"><div class="Beacon"><div class="sc-AxheI gkgGqG"></div><div class="sc-AxhCb ljOwHR BeaconContainer is-configDisplayRight is-manual"><button data-cy="beacon-close-button" class="sc-AxhUy iLeqUU c-BeaconCloseButton is-rendered"><span class="sc-AxgMl lbFAFy c-BeaconCloseButton__label"><span data-cy="Text" class="Textcss__TextUI-cdrlb6-0 eMuyYC c-Text is-span is-13 is-default">Close</span></span><span aria-hidden="true" role="img" data-cy="Icon" class="Iconcss__IconUI-vne9x2-0 lhTUZR c-Icon is-noInteract is-iconName-cross-small is-24 c-BeaconCloseButton__iconClose" data-icon-name="cross-small" title="cross-small"><span class="c-Icon__icon"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill-rule="evenodd" d="M13.414 12l2.293-2.293a.999.999 0 1 0-1.414-1.414L12 10.586 9.707 8.293a.999.999 0 1 0-1.414 1.414L10.586 12l-2.293 2.293a.999.999 0 1 0 1.414 1.414L12 13.414l2.293 2.293a.997.997 0 0 0 1.414 0 .999.999 0 0 0 0-1.414L13.414 12z"></path></svg></span><span data-cy="VisuallyHidden" class="VisuallyHiddencss__VisuallyHiddenUI-f040fk-0 btmtsE c-VisuallyHidden">cross-small</span></span></button></div></div></div></div></body></html>`))
	http.DefaultServeMux.HandleFunc("/gz", f.HttpHandler)
	http.ListenAndServe("0.0.0.0:12346", nil)
}
