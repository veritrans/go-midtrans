CreditCardField.prototype.constructor = CreditCardField;

/**
 * This library is basically very essential version of CardJs. 
 * It has been decoupled from the original implementation so that it does not
 * overly magical, and so that it only support the credit card field,
 * nothing else. It has to be used along with jQuery
 *
 * @author Adam Pahlevi Baihaqi
 * @author Colin Stannard
 */
 
function CreditCardField(elem) {
  this.elem = jQuery(elem)
  this.initCardNumberInput()
  this.elem.empty()
  this.setupCardNumberInput()
  this.refreshCreditCardTypeIcon()
}

CreditCardField.KEYS = {
  "0": 48,
  "9": 57,
  "DELETE": 46,
  "BACKSPACE": 8,
  "ARROW_LEFT": 37,
  "ARROW_RIGHT": 39,
  "ARROW_UP": 38,
  "ARROW_DOWN": 40,
  "HOME": 36,
  "END": 35,
  "TAB": 9,
  "A": 65,
  "X": 88,
  "C": 67,
  "V": 86
}

CreditCardField.CREDIT_CARD_NUMBER_MASK = "XXXX XXXX XXXX XXXX"
CreditCardField.CREDIT_CARD_NUMBER_PLACEHOLDER = "Card number"

CreditCardField.prototype.initCardNumberInput = function() {
  var $this = this

  this.cardNumberInput = this.elem
  this.cardNumberInput.attr("type", "tel")
  if(!this.cardNumberInput.attr("placeholder")) {
    this.cardNumberInput.attr("placeholder", CreditCardField.CREDIT_CARD_NUMBER_PLACEHOLDER)
  }
  this.cardNumberInput.attr("maxlength", CreditCardField.CREDIT_CARD_NUMBER_MASK.length)
  this.cardNumberInput.attr("x-autocompletetype", "cc-number")
  this.cardNumberInput.attr("autocompletetype", "cc-number")
  this.cardNumberInput.attr("autocorrect", "off")
  this.cardNumberInput.attr("spellcheck", "off")
  this.cardNumberInput.attr("autocapitalize", "off")
}

CreditCardField.prototype.setupCardNumberInput = function() {
  $this = this

  new_element = jQuery("<div class='card-number-wrapper'></div>")

  // clone the CC field, and give it callbacks
  // cloning is needed because it will be replaced, otherwise
  // it cannot be replaced because will be treated as parent
  clonedCcField = $(this.elem).clone()
  clonedCcField.keydown(CreditCardField.handleCreditCardNumberKey)
  clonedCcField.keyup(function(e) {
    $this.refreshCreditCardTypeIcon()
  })
  clonedCcField.change(CreditCardField.handleCreditCardNumberChange)
  this.cardNumberInput = clonedCcField
  new_element.append(clonedCcField)

  new_element.append("<div class='card-type-icon'></div>")
  new_element.append("<div class='icon'></div>")

  parent_div = jQuery("<div class='credit-card-field'></div>")
  parent_div.append(new_element)

  this.elem.replaceWith(parent_div)
  this.elem = parent_div
}

CreditCardField.keyCodeFromEvent = function(e) {
  return e.which || e.keyCode
}

CreditCardField.keyIsNumber = function(event) {
  var keyCode = CreditCardField.keyCodeFromEvent(event)
  return keyCode >= CreditCardField.KEYS["0"] && 
    keyCode <= CreditCardField.KEYS["9"]
}

CreditCardField.keyIsDelete = function(event) {
  return CreditCardField.keyCodeFromEvent(event) == 
    CreditCardField.KEYS["DELETE"]
}

CreditCardField.keyIsBackspace = function(event) {
  return CreditCardField.keyCodeFromEvent(event) ==
    CreditCardField.KEYS["BACKSPACE"]
}

CreditCardField.filterNumberOnlyKey = function(e) {
  var isNumber = CreditCardField.keyIsNumber(e)

  var isDeletion = function(e) {
    return CreditCardField.keyIsDelete(e) || CreditCardField.keyIsBackspace(e)
  }(e) 

  var isArrow = function(e) {
    var keyCode = CreditCardField.keyCodeFromEvent(e)
    return keyCode >= CreditCardField.KEYS["ARROW_LEFT"] &&
      keyCode <= CreditCardField.KEYS["ARROW_DOWN"]
  }(e)

  var isNavigation = function(e) {
    var keyCode = CreditCardField.keyCodeFromEvent(e)
    return keyCode == CreditCardField.KEYS["HOME"] ||
      keyCode == CreditCardField.KEYS["END"]
  }(e)

  var isKeyboardCommand = function(e) {
    var keyCode = CreditCardField.keyCodeFromEvent(e)
    return (e.ctrlKey || e.metaKey) && (
        keyCode == CreditCardField.KEYS["A"] ||
        keyCode == CreditCardField.KEYS["X"] ||
        keyCode == CreditCardField.KEYS["C"] ||
        keyCode == CreditCardField.KEYS["V"]
      )
  }(e) 
  
  var isTab = function(e) {
    return CreditCardField.keyCodeFromEvent(e) == CreditCardField.KEYS["TAB"]
  }(e) 

  if (!isNumber && !isDeletion && !isArrow && !isNavigation && !isKeyboardCommand && !isTab) {
    e.preventDefault()
  }
}

CreditCardField.handleMaskedNumberInputKey = function(e, mask) {
  CreditCardField.filterNumberOnlyKey(e)

  var keyCode = CreditCardField.keyCodeFromEvent(e)
  var digit = String.fromCharCode(keyCode)
  var element = e.target
  
  var caretStart = function caretStartPosition(element) {
    if (typeof element.selectionStart == "number") {
      return element.selectionStart
    }
    return false
  }(element)

  var caretEnd = function caretEndPosition(element) {
    if (typeof element.selectionEnd == "number") {
      return element.selectionEnd
    }
    return false
  }(element)

  // calculate normalised caret position
  var normalisedStartCaretPosition = CreditCardField
    .normaliseCaretPosition(mask, caretStart) 
  var normalisedEndCaretPosition = CreditCardField
    .normaliseCaretPosition(mask, caretEnd)

  var newCaretPosition = caretStart

  var isNumber = CreditCardField.keyIsNumber(e)
  var isDelete = CreditCardField.keyIsDelete(e)
  var isBackspace = CreditCardField.keyIsBackspace(e)

  if (isNumber || isDelete || isBackspace) {
    e.preventDefault()
    var rawText = $(element).val()
    var numbersOnly = CreditCardField.numbersOnlyString(rawText)

    var rangeHighlighted = normalisedEndCaretPosition > normalisedStartCaretPosition

    // Remove values highlighted (if highlighted)
    if (rangeHighlighted) {
      numbersOnly = (numbersOnly.slice(0, normalisedStartCaretPosition) +
        numbersOnly.slice(normalisedEndCaretPosition))
    }

    // Forward Action
    if (caretStart != mask.length) {
      // insert number digit
      if (isNumber && rawText.length < mask.length) {
        numbersOnly = (numbersOnly.slice(0, normalisedStartCaretPosition) + 
          digit + (numbersOnly.slice(normalisedStartCaretPosition)))
        newCaretPosition = Math.max(
          CreditCardField.denormaliseCaretPosition(mask, normalisedStartCaretPosition + 1),
          CreditCardField.denormaliseCaretPosition(mask, normalisedStartCaretPosition + 2) - 1
        )
      }

      // delete
      if (isDelete) {
        numbersOnly = (numbersOnly.slice(0, normalisedStartCaretPosition) + 
          numbersOnly.slice(normalisedStartCaretPosition + 1))
      }
    }

    // Backward Action
    if (caretStart != 0) {
      if (isBackspace && !rangeHighlighted) {
        numbersOnly = (numbersOnly.slice(0, normalisedStartCaretPosition - 1) +
          numbersOnly.slice(normalisedStartCaretPosition))
        newCaretPosition = CreditCardField.denormaliseCaretPosition(mask, 
          normalisedStartCaretPosition - 1)
      }
    }

    $(element).val(CreditCardField.applyFormatMask(numbersOnly, mask))
    CreditCardField.setCaretPosition(element, newCaretPosition)
  }
}

/* 
 * Set the caret position of the given element
 *
 * @param element
 * @param caretPos
 */
CreditCardField.setCaretPosition = function(element, caretPos) {
  if(element != null) {
    if (element.createTextRange) {
      var range = element.createTextRange()
      range.move('character', caretPos)
      range.select()
    } else {
      if (element.selectionStart) {
        element.focus()
        element.setSelectionRange(caretPos, caretPos)
      } else { element.focus() }
    }
  }
}

CreditCardField.applyFormatMask = function(string, mask) {
  var formattedString = ""
  var numberPos = 0

  for(var j = 0; j < mask.length; j++) {
    var currentMaskChar = mask[j]
    if (currentMaskChar == "X") {
      var digit = string.charAt(numberPos)
      if (!digit) { break; }
      formattedString += string.charAt(numberPos)
      numberPos++
    } else {
      formattedString += currentMaskChar
    }
  }
  return formattedString
}

/* 
 * Normalise the caret position for the given mask.
 *
 * @param mask
 * @param caretPosition
 * @returns {number}
 */
CreditCardField.normaliseCaretPosition = function(mask, caretPosition) {
  var numberPos = 0
  if (caretPosition < 0 || caretPosition > mask.length) { return 0 }
  for(var i = 0; i < mask.length; i++) {
    if(i == caretPosition) { return numberPos }
    if(mask[i] == "X") { numberPos++ }
  }
  return numberPos
}

/**
 * Denormalise the caret position for the given mask
 *
 * @param mask
 * @param caretPosition
 * @returns {*}
 */
CreditCardField.denormaliseCaretPosition = function(mask, caretPosition) {
  var numberPos = 0
  if(caretPosition < 0 || caretPosition > mask.length) { return 0 }
  for(var i = 0; i < mask.length; i++) {
    if(numberPos == caretPosition) { return i }
    if(mask[i] == "X") { numberPos++ }
  }
  return mask.length
}

CreditCardField.handleCreditCardNumberKey = function(e) {
  CreditCardField.handleMaskedNumberInputKey(e, CreditCardField.CREDIT_CARD_NUMBER_MASK)
}

CreditCardField.numbersOnlyString = function(string) {
  var numbersOnlyString = ""
  for(var i = 0; i < string.length; i++) {
    var currentChar = string.charAt(i);
    var isValid = !isNaN(parseInt(currentChar))
    if (isValid) { numbersOnlyString += currentChar }
  }
  return numbersOnlyString
}

CreditCardField.prototype.refreshCreditCardTypeIcon = function() {
  var ccVal = CreditCardField.numbersOnlyString(this.cardNumberInput.val())

  var cardType = function(number) {
    // Visa
    var re = new RegExp("^4")
    if (number.match(re) != null)
      return "Visa";

    // Mastercard
    re = new RegExp("^5[1-5]")
    if (number.match(re) != null)
      return "Mastercard"

    // AMEX
    re = new RegExp("^3[47]")
    if (number.match(re) != null)
      return "AMEX"
    
    // Discover
    re = new RegExp("^(6011|622(12[6-9]|1[3-9][0-9]|[2-8][0-9]{2}|9[0-1][0-9]|92[0-5]|64[4-9])|65)")
    if (number.match(re) != null)
      return "Discover"

    // Dinners
    re = new RegExp("^36")
    if (number.match(re) != null)
      return "Diners"

    // Dinners - Carte Blanche
    re = new RegExp("^30[0-5]")
    if (number.match(re) != null)
      return "Diners - Carte Blanche"

    // JCB
    re = new RegExp("^35(2[89]|[3-8][0-9])")
    if (number.match(re) != null)
      return "JCB"

    // Visa Electron
    re = new RegExp("^(4026|417500|4508|4844|491(3|7))")
    if (number.match(re) != null)
      return "Visa Electron"

    return ""
  }(ccVal)

  var cardTypeIcon = this.elem.find(".card-type-icon")

  switch(cardType) {
    case "Visa Electron":
    case "Visa":
      cardTypeIcon.attr("class", "card-type-icon show visa")
      break;
    case "Mastercard":
      cardTypeIcon.attr("class", "card-type-icon show master-card")
      break;
    case "AMEX":
      cardTypeIcon.attr("class", "card-type-icon show american-express")
      break;
    case "Discover":
      cardTypeIcon.attr("class", "card-type-icon show discover")
      break;
    case "Diners - Carte Blanche":
    case "Diners": 
      cardTypeIcon.attr("class", "card-type-icon show diners")
      break;
    case "JCB":
      cardTypeIcon.attr("class", "card-type-icon show jcb")
      break;
    default:
      this.elem.find(".card-number-wrapper .card-type-icon")
        .removeClass("show")
  }
}

// initialize to jQuery
$(document).ready(function() {
  var methods = {
    init: function() { new CreditCardField(this) }
  }
  $.fn.makeCreditCardField = function(methodOrOptions) {
    if (methods[methodOrOptions]) {
      return methods[methodOrOptions].apply(this, 
                                            Array.prototype.slice.call(arguments, 1))
    } else if (typeof methodOrOptions === "object" || !methodOrOptions) {
      return methods.init.apply(this, arguments)
    } else {
      $.error("Method " + methodOrOptions + " does not exist on")
    }
  }
})