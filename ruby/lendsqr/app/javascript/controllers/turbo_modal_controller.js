import { Controller } from "@hotwired/stimulus"

// Connects to data-controller="turbo-modal"
export default class extends Controller {
  hideModal() {
    this.element.parentElement.removeAttribute("src");
    this.element.remove();
  }

  closeBackground(e) {
    if (e && Array.from(this.element.querySelectorAll('*')).find(node => node.isEqualNode(e.target))) {
      return
    }
    this.hideModal()
  }
}
