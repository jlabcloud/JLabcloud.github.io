        angular.module('jlabsApp', [])
            .controller('DropdownController', function() {
                this.menus = {};
                this.toggleMenu = function(menuId) {
                    this.menus[menuId] = !this.menus[menuId];
                };
                this.isMenuOpen = function(menuId) {
                    return !!this.menus[menuId];
                };
            })
            .controller('SearchController', function() {
                this.query = '';
                this.handleKeyPress = function(event) {
                    if (event.key === 'Enter') {
                        console.log('Search query:', this.query);
                        // Add your search logic here
                    }
                };
            });
