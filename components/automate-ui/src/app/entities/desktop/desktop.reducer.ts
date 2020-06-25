import { set, pipe, find, concat, remove } from 'lodash/fp';
import { EntityStatus } from '../entities';
import { DesktopActionTypes, DesktopActions } from './desktop.actions';
import {
  DailyCheckInCountCollection,
  TopErrorsCollection,
  CountedDurationCollection,
  NodeMetadataCount,
  NodeMetadataCountValue,
  NodeMetadataCountType,
  DailyNodeRuns,
  Desktop,
  Filter,
  TermFilter,
  SortOrder,
  Selected,
  Terms
} from './desktop.model';

export interface DesktopEntityState {
  dailyCheckInCountCollection: DailyCheckInCountCollection;
  getDailyCheckInTimeSeriesStatus: EntityStatus;
  selected: Selected;
  topErrorCollection: TopErrorsCollection;
  getTopErrorCollectionStatus: EntityStatus;
  unknownDesktopDurationCounts: CountedDurationCollection;
  getUnknownDesktopDurationCountsStatus: EntityStatus;
  nodeMetadataCounts: NodeMetadataCount[];
  getNodeMetadataCountsStatus: EntityStatus;
  dailyNodeRuns: DailyNodeRuns;
  desktopListTitle: string;
  desktops: Desktop[];
  getDesktopsStatus: EntityStatus;
  desktopsTotal: number;
  getDesktopsTotalStatus: EntityStatus;
  getDesktopsFilter: Filter;
}

export const desktopEntityInitialState: DesktopEntityState = {
  dailyCheckInCountCollection: { buckets: [], updated: new Date(0)},
  getDailyCheckInTimeSeriesStatus: EntityStatus.notLoaded,
  selected: {
    desktop: undefined,
    daysAgo: 3
  },
  topErrorCollection: {items: [], updated: new Date(0)},
  getTopErrorCollectionStatus: EntityStatus.notLoaded,
  unknownDesktopDurationCounts: {items: [], updated: new Date(0)},
  getUnknownDesktopDurationCountsStatus: EntityStatus.notLoaded,
  nodeMetadataCounts: [],
  getNodeMetadataCountsStatus: EntityStatus.notLoaded,
  desktopListTitle: 'Desktops',
  desktops: [],
  getDesktopsStatus: EntityStatus.notLoaded,
  desktopsTotal: 0,
  getDesktopsTotalStatus: EntityStatus.notLoaded,
  dailyNodeRuns: {
    durations: { buckets: [], updated: new Date(0)},
    daysAgo: 14,
    nodeId: '',
    status: EntityStatus.notLoaded
  },
  getDesktopsFilter: {
    currentPage: 1,
    pageSize: 10,
    sortingField: Terms.DesktopName,
    sortingOrder: SortOrder.Ascending,
    terms: []
  }
};

export function desktopEntityReducer(state: DesktopEntityState = desktopEntityInitialState,
  action: DesktopActions): DesktopEntityState {

  switch (action.type) {
    case DesktopActionTypes.SET_SELECTED_DESKTOP:
      return set('selected.desktop', action.payload.desktop, state);

    case DesktopActionTypes.SET_SELECTED_DAYS_AGO:
      return set('selected.daysAgo', action.payload.daysAgo, state);

    case DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES:
      return set('getDailyCheckInTimeSeriesStatus', EntityStatus.loading, state);

    case DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES_SUCCESS:
      return pipe(
        set('getDailyCheckInTimeSeriesStatus', EntityStatus.loadingSuccess),
        set('dailyCheckInCountCollection', action.payload))(state) as DesktopEntityState;

    case DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES_FAILURE:
      return set('getDailyCheckInTimeSeriesStatus', EntityStatus.loadingFailure, state);

    case DesktopActionTypes.GET_DAILY_NODE_RUNS_STATUS_TIME_SERIES:
      return pipe(
        set('dailyNodeRuns.status', EntityStatus.loading),
        set('dailyNodeRuns.nodeId', action.nodeId),
        set('dailyNodeRuns.daysAgo', action.daysAgo))(state) as DesktopEntityState;

    case DesktopActionTypes.GET_DAILY_NODE_RUNS_STATUS_TIME_SERIES_SUCCESS:
      return pipe(
        set('dailyNodeRuns.status', EntityStatus.loadingSuccess),
        set('dailyNodeRuns.durations', action.payload))(state) as DesktopEntityState;

    case DesktopActionTypes.GET_DAILY_NODE_RUNS_STATUS_TIME_SERIES_FAILURE:
      return set('dailyNodeRuns.status', EntityStatus.loadingFailure, state);

    case DesktopActionTypes.GET_TOP_ERRORS_COLLECTION:
      return set('getTopErrorCollectionStatus', EntityStatus.loading, state);

    case DesktopActionTypes.GET_TOP_ERRORS_COLLECTION_SUCCESS:
      return pipe(
        set('getTopErrorCollectionStatus', EntityStatus.loadingSuccess),
        set('topErrorCollection', action.payload))(state) as DesktopEntityState;

    case DesktopActionTypes.GET_TOP_ERRORS_COLLECTION_FAILURE:
      return set('getTopErrorCollectionStatus', EntityStatus.loadingFailure, state);

    case DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS:
      return set('getUnknownDesktopDurationCountsStatus', EntityStatus.loading, state);

    case DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS_SUCCESS:
      return pipe(
        set('getUnknownDesktopDurationCountsStatus', EntityStatus.loadingSuccess),
        set('unknownDesktopDurationCounts', action.payload))(state) as DesktopEntityState;

    case DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS_FAILURE:
      return set('getUnknownDesktopDurationCountsStatus', EntityStatus.loadingFailure, state);

    case DesktopActionTypes.GET_NODE_METADATA_COUNTS:
      return set('getNodeMetadataCountsStatus', EntityStatus.loading, state);

    case DesktopActionTypes.GET_NODE_METADATA_COUNTS_SUCCESS:
      return pipe(
        set('getNodeMetadataCountsStatus', EntityStatus.loadingSuccess),
        set('nodeMetadataCounts', nodeMetadataCounts(action.payload, state.getDesktopsFilter.terms))
      )(state) as DesktopEntityState;

    case DesktopActionTypes.GET_NODE_METADATA_COUNTS_FAILURE:
      return set('getNodeMetadataCountsStatus', EntityStatus.loadingFailure, state);

    case DesktopActionTypes.UPDATE_DESKTOP_LIST_TITLE:
      return set('desktopListTitle', action.payload)(state);

    case DesktopActionTypes.GET_DESKTOPS:
      return set('getDesktopsStatus', EntityStatus.loading, state);

    case DesktopActionTypes.GET_DESKTOPS_SUCCESS:
      return pipe(
        set('getDesktopsStatus', EntityStatus.loadingSuccess),
        set('desktops', action.payload))(state) as DesktopEntityState;

    case DesktopActionTypes.GET_DESKTOPS_FAILURE:
      return set('getDesktopsStatus', EntityStatus.loadingFailure, state);

    case DesktopActionTypes.GET_DESKTOPS_TOTAL:
      return set('getDesktopsTotalStatus', EntityStatus.loading, state);

    case DesktopActionTypes.GET_DESKTOPS_TOTAL_SUCCESS:
      return pipe(
        set('getDesktopsTotalStatus', EntityStatus.loadingSuccess),
        set('desktopsTotal', action.payload))(state) as DesktopEntityState;

    case DesktopActionTypes.GET_DESKTOPS_TOTAL_FAILURE:
      return set('getDesktopsTotalStatus', EntityStatus.loadingFailure, state);

    case DesktopActionTypes.UPDATE_DESKTOPS_FILTER_CURRENT_PAGE:
      return set('getDesktopsFilter.currentPage', action.payload.page)(state);

    case DesktopActionTypes.ADD_DESKTOPS_FILTER_TERM: {
      const terms = concat(state.getDesktopsFilter.terms, [action.payload.term]);
      return pipe(
        set('getDesktopsFilter.terms', terms),
        set('getDesktopsFilter.currentPage', 1),
        set('nodeMetadataCounts', nodeMetadataCounts(state.nodeMetadataCounts, terms))
      )(state) as DesktopEntityState;
    }

    case DesktopActionTypes.UPDATE_DESKTOPS_FILTER_TERMS:
      return pipe(
        set('getDesktopsFilter.terms', action.payload.terms ),
        set('getDesktopsFilter.currentPage', 1),
        set('nodeMetadataCounts',
          nodeMetadataCounts(state.nodeMetadataCounts, action.payload.terms))
      )(state) as DesktopEntityState;

    case DesktopActionTypes.REMOVE_DESKTOPS_FILTER_TERM: {
      const terms = remove<TermFilter>(
        term => term.type === action.payload.term.type,
        state.getDesktopsFilter.terms);
      return pipe(
        set('getDesktopsFilter.terms', terms),
        set('getDesktopsFilter.currentPage', 1),
        set('nodeMetadataCounts', nodeMetadataCounts(state.nodeMetadataCounts, terms))
      )(state) as DesktopEntityState;
    }

    case DesktopActionTypes.UPDATE_DESKTOPS_SORT_TERM:
      let order = SortOrder.Ascending;
      if (action.payload.term === state.getDesktopsFilter.sortingField) {
        if (state.getDesktopsFilter.sortingOrder === SortOrder.Ascending) {
          order = SortOrder.Descending;
        }
      }

      return pipe(
        set('getDesktopsFilter.sortingField', action.payload.term),
        set('getDesktopsFilter.sortingOrder', order)
      )(state) as DesktopEntityState;

    case DesktopActionTypes.UPDATE_DESKTOPS_DATE_TERM:
      return pipe(
        set('getDesktopsFilter.start', action.payload.start),
        set('getDesktopsFilter.end', action.payload.end)
      )(state) as DesktopEntityState;

    default:
      return state;

  }
}

function nodeMetadataCounts(counts: NodeMetadataCount[], terms: TermFilter[]): NodeMetadataCount[] {
  const maxCount = 3;

  return counts.map((count: NodeMetadataCount) => {
    const label = nodeMetadataCountLabel(count, maxCount);
    const values = count.values.map((value: NodeMetadataCountValue) => {
      const typeIsSelected = !!find({ type: count.type }, terms);
      const valueIsSelected = !!find({ value: value.value }, terms);
      const disabled = typeIsSelected ? !valueIsSelected : false;
      const checked = typeIsSelected && valueIsSelected;
      return { ...value, disabled, checked };
    }).slice(0, maxCount);

    return { ...count, label, values };
  });
}

function nodeMetadataCountLabel(count: NodeMetadataCount, maxCount: number): string {
  switch (count.type) {
    case NodeMetadataCountType.Domain:
      return `Top ${maxCount} Domains`;
    case NodeMetadataCountType.Platform:
      return `Top ${maxCount} Platforms`;
    case NodeMetadataCountType.Environment:
      return `Top ${maxCount} Environments`;
    case NodeMetadataCountType.Status:
      return 'Last Run';
  }
}
