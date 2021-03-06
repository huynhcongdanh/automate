import { Component, EventEmitter, Input, OnChanges, Output, HostBinding } from '@angular/core';
import { MatOptionSelectionChange } from '@angular/material/core/option';
import { PageSizeChangeEvent } from 'app/entities/desktop/desktop.model';


@Component({
  selector: 'app-page-picker',
  templateUrl: './page-picker.component.html',
  styleUrls: ['./page-picker.component.scss']
})
export class PagePickerComponent implements OnChanges {

  @Input() total: number;
  @Input() perPage: number;
  @Input() page: number;
  @Input() maxPageItems = 10;
  @Input() @HostBinding('class.forDesktop') forDesktop = false;
  @Input() @HostBinding('class.fullScreened') fullScreened = false;

  @Output() pageChanged = new EventEmitter<number>();
  @Output() pageSizeChanged = new EventEmitter<PageSizeChangeEvent>();

  itemStartCount: number;
  itemEndCount: number;
  allPages: number[];

  perPageOptions = [10, 20, 50];
  selectablePages = [];

  // Page number variables
  first = 1;
  prev = 1;
  next = 1;
  last = 1;

  private getItemStartCount() {
    this.itemStartCount = this.page === 1 ? 1 : ((this.page - 1) * this.perPage + 1);
  }

  private getItemEndCount() {
    this.itemEndCount = this.page === this.last ? this.total : this.page * this.perPage;
  }

  private getAllPages() {
    const pageCount = Math.ceil(this.total / this.perPage);
    this.allPages = Array(pageCount).fill(0).map((_x, i) => i + 1);
  }

  private setCounts() {
    this.getItemStartCount();
    this.getItemEndCount();
    this.getAllPages();
  }

  ngOnChanges() {
    this.last = Math.ceil(this.total / this.perPage) || 1;
    this.prev = (this.page === this.first) ? null : this.page - 1;
    this.next = (this.page === this.last) ? null : this.page + 1;

    const pages = [];
    const selectedIndex = ((this.page - 1) % this.maxPageItems);

    for (let pageIndex = selectedIndex; pageIndex >= 0; --pageIndex) {
      const prev = this.page + (pageIndex - selectedIndex);
      pages[pageIndex] = prev;
    }

    pages[selectedIndex] = this.page;

    for (let pageIndex = selectedIndex; pageIndex < this.maxPageItems; ++pageIndex) {
      const next = this.page + (pageIndex - selectedIndex);
      if (next <= this.last) {
        pages[pageIndex] = next;
      }
    }
    this.selectablePages = pages;

    if ( this.forDesktop ) {
      this.setCounts();
    }
  }

  onItemTap(value) {
    this.pageChanged.emit(value);
  }

  isNull(value) {
    return value === null;
  }

  trackBy(_index, item) {
    return item;
  }

  handleSelectItem(event: MatOptionSelectionChange, value: number): void {
    if (event.isUserInput) {
      this.pageChanged.emit(value);
    }
  }

  handleSelectPerPageItems(event: MatOptionSelectionChange, value: number): void {
    if (event.isUserInput) {
      const updatedPageNumber = Math.ceil(this.itemStartCount / value);

      this.pageSizeChanged.emit({
        pageSize: value,
        updatedPageNumber
      });
    }
  }
}
