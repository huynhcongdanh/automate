import { Component, EventEmitter, Input, Output } from '@angular/core';
import { cloneDeep } from 'lodash/fp';
import { ProjectsFilterOption } from 'app/services/projects-filter/projects-filter.reducer';
import { ProjectConstants } from 'app/entities/projects/project.model';

const { ALL_PROJECTS_LABEL } = ProjectConstants;

@Component({
  selector: 'app-projects-filter-dropdown',
  templateUrl: './projects-filter-dropdown.component.html',
  styleUrls: [ './projects-filter-dropdown.component.scss' ]
})
export class ProjectsFilterDropdownComponent {
  @Input() options: ProjectsFilterOption[] = [];

  @Input() selectionLabel = ALL_PROJECTS_LABEL;

  @Input() selectionCount = 0;

  @Input() selectionCountVisible = false;

  @Input() selectionCountActive = false;

  @Input() dropdownCaretVisible = false;

  @Input() filterVisible = false;

  @Output() onSelection = new EventEmitter<ProjectsFilterOption[]>();
  @Output() onOptionChange = new EventEmitter<ProjectsFilterOption[]>();

  editableOptions: ProjectsFilterOption[] = [];
  // Filtered options is merely a copy of the editable options
  // so they can be filtered while maintaining the actual options.
  filteredOptions: ProjectsFilterOption[] = [];

  optionsEdited = false;

  dropdownActive = false;

  resetOptions() {
    if (!this.optionsEdited) {
      this.filteredOptions = this.editableOptions = cloneDeep(this.options);
    }
  }

  get filteredSelectedCount(): string {
    const count = this.filteredOptions.filter(option => option.checked).length;
    return count > 99 ? '99+' : count.toString();
  }

  handleFilterKeyUp(filterValue: string): void {
    this.filteredOptions = this.editableOptions
      .filter(option => option.label.toLowerCase().indexOf(filterValue.toLowerCase()) > -1);
  }

  handleLabelClick() {
    this.resetOptions();
    if (this.editableOptions.length > 1) {
      this.dropdownActive = !this.dropdownActive;
    }
  }

  handleEscape(): void {
    this.optionsEdited = false;
    this.resetOptions();
    this.dropdownActive = false;
    this.onOptionChange.emit(this.editableOptions);
  }

  handleOptionChange(event, label) {
    // sets the new state of the specific checkbox
    this.editableOptions
      .find(option => option.label === label).checked = event.detail;
    this.optionsEdited = true;
  }

  handleApplySelection() {
    this.dropdownActive = false;
    this.optionsEdited = false;
    this.onOptionChange.emit(this.editableOptions);
    this.onSelection.emit(this.editableOptions);
  }

  handleClearSelection() {
    // TODO: Should ideally set to true only when some projects were selected upon opening
    this.optionsEdited = true; // mark as edited
    this.editableOptions.map(option => option.checked = false); // uncheck all the options
  }

  handleArrowUp(event: KeyboardEvent) {
    event.preventDefault();

    const element = (event.target as Element).previousElementSibling;
    if (element) {
      (element as HTMLElement).focus();
    }
  }

  handleArrowDown(event: KeyboardEvent) {
    event.preventDefault();

    const element = (event.target as Element).nextElementSibling;
    if (element) {
      (element as HTMLElement).focus();
    }
  }
}
