/**
 * 
 * Manipulating the DOM exercise.
 * Exercise programmatically builds navigation,
 * scrolls to anchors from navigation,
 * and highlights section in viewport upon scrolling.
 * 
 * Dependencies: None
 * 
 * JS Version: ES2015/ES6
 * 
 * JS Standard: ESlint
 * 
*/

/**
 * Comments should be present at the beginning of each procedure and class.
 * Great to have comments before crucial code sections within the procedure.
*/

/**
 * Define Global Variables
 * 
*/

// Create an empty to store the sections
let sections = [];
// current active class
const ACTIVE_CLASS_SECTION = 'your-active-class';
// Menu item active class
const ACTIVE_CLASS_MENU = 'menu-active-class';

/**
 * End Global Variables
*/



/**
 * Start Helper Functions
 * 
*/

/**
 * @description Set property isActive to true for element with id equal to active_section
 *
 * @param {array} rows
 *        example
 *        [ {"id":"section1","nav":"Section 1","isActive":false},
 *          {"id":"section2","nav":"Section 2","isActive":false},
 *          {"id":"section3","nav":"Section 3","isActive":false},
 *          {"id":"section4","nav":"Section 4","isActive":false},
 *          {"id":"section5","nav":"Section 5","isActive":false}
 *        ]
 * @param {string} active_section
 *        example
 *        section 5
 * @returns {array} the list with the updated state
 *        example
 *        [ {"id":"section1","nav":"Section 1","isActive":false},
 *          {"id":"section2","nav":"Section 2","isActive":false},
 *          {"id":"section3","nav":"Section 3","isActive":false},
 *          {"id":"section4","nav":"Section 4","isActive":false},
 *          {"id":"section5","nav":"Section 5","isActive":true}
 *        ]
*/
function updateSection(rows, active_section) {
   return rows.map(row => {
      const found = row.id === active_section;
      if (row.isActive !== found) {
         row.isActive = found;
      }
      return row;
   })
}

/**
 * @description Update DOM base on rows
 *
 * @param {array} rows
 *        example
 *        [ {"id":"section1","nav":"Section 1","isActive":true},
 *          {"id":"section2","nav":"Section 2","isActive":false},
 *          {"id":"section3","nav":"Section 3","isActive":false},
 *          {"id":"section4","nav":"Section 4","isActive":false},
 *          {"id":"section5","nav":"Section 5","isActive":false}
 *        ]
 * @param {string} active_section
 *        example
 *        your-active-class
 * @param {string} active_menu
 *        example
 *        menu-active-class
*/
function updateDOM(rows, active_section, active_menu) {

   rows.forEach(item => {
      const sectionBlk = document.getElementById(item.id);
      const menuLink = document.getElementById(`link-${item.id}`);

      if (item.isActive) {
         // if section is active and don't have active class then add active class to section and menu classlist
         if (!sectionBlk.classList.contains(active_section)) {
            sectionBlk.classList.add(active_section);
            menuLink.classList.add(active_menu);
         }
      } else {
         // if section is not active and has active class then remove the active class to section and menu classlist
         if (sectionBlk.classList.contains(active_section)) {
           sectionBlk.classList.remove(active_section);
           menuLink.classList.remove(active_menu);
         }
      }
   })
}

function isInViewPort(element) {
   const rect = element.getBoundingClientRect();
   return rect.top < window.innerHeight / 2 && rect.bottom >= window.innerHeight / 2;
};

function scrollHandler () {

   const all_sections = document.querySelectorAll('section');
   all_sections.forEach(element => {
      const id = element.getAttribute('id');
      const nav = element.getAttribute('data-nav');
      if (isInViewPort(element)) {
          // Update the data structure
          sections = updateSection(sections, id);
          // Update the DOM
          updateDOM(sections, ACTIVE_CLASS_SECTION, ACTIVE_CLASS_MENU);
      }

   });
}


/**
 * End Helper Functions
*/



/**
 * Begin Main Functions
 * 
*/

function main() {

   // Select all section elements in html
   const all_sections = document.querySelectorAll('section');

   // Select navbar__list element in html
   const list_navbar = document.querySelector('#navbar__list');

   // initialize sections
   all_sections.forEach(element => {
      const id = element.getAttribute('id');
      const nav = element.getAttribute('data-nav');

      const is_active = element.classList.contains('active');
      sections.push({ id, nav, is_active });
   });

   // build navbar
   sections.forEach(section => {
      // Create html list element
      const element_list = document.createElement('li');

      // add navbar__menu to element_list classList
      element_list.classList.add('navbar__menu');

      // build anchor element with class menu__link, id, href and content
      const element_link = document.createElement('a');
      element_link.textContent = section.nav;
      element_link.setAttribute('id', `link-${section.id}`);
      element_link.setAttribute('href', `#${section.id}`);
      element_link.classList.add('menu__link');
      if (list_navbar.childElementCount === 0) {
         element_link.classList.add('menu-active-class');
      }

      // Add the built anchor element to element_list as a child
      element_list.appendChild(element_link);
      // finally add element_list to the navbar element
      list_navbar.appendChild(element_list);
   });

   // Add class 'active' to section when near top of viewport
   // https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API#creating_an_intersection_observer
   /*
   const options = {
     root: null,
     rootMargin: '0px',
     threshold: 0.5,
   };

   const observer = new IntersectionObserver(entries => {
      entries.forEach(entry => {
         if (entry.isIntersecting) {
            // Update the data structure
            sections = updateSection(sections, entry.target.id);
            // Update the DOM
            updateDOM(sections, ACTIVE_CLASS_SECTION, ACTIVE_CLASS_MENU);
         }
      });
   }, options);

   sections.forEach(section => {
      observer.observe(document.getElementById(section.id));
   });
   */

   sections.forEach(section => {
      if (isInViewPort(document.getElementById(section.id))) {
         // Update the data structure
         sections = updateSection(sections, section.id);
         // Update the DOM
         updateDOM(sections, ACTIVE_CLASS_SECTION, ACTIVE_CLASS_MENU);
      }
   });

   // scroll view operation
   const links = document.querySelectorAll('.menu__link');
   links.forEach(link => {
      // register eventListener on click by smooth scroll into the active view
      link.addEventListener('click', (event) => {
         event.preventDefault();
         const sectionId = link.getAttribute('href').substring(1);
         sections = updateSection(sections, sectionId);
         updateDOM(sections, ACTIVE_CLASS_SECTION, ACTIVE_CLASS_MENU);
         document.getElementById(sectionId).scrollIntoView({
            behavior: 'smooth',
         });
      });
   });

   window.addEventListener('scroll', scrollHandler);

}

/**
 * End Main Functions
*/



/**
 * Begin Events
 * 
*/

// Build menu
document.addEventListener('DOMContentLoaded', () => {
  main();
});

/**
 * End Events
*/
