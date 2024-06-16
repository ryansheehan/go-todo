import './style.css'
import htmx from 'htmx.org';

declare global {
  interface Window {
    htmx: unknown
  }
}

window.htmx = htmx;

console.log('hello world!')
