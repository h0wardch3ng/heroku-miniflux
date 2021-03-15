// Code generated by go generate; DO NOT EDIT.

package template // import "miniflux.app/template"

var templateCommonMap = map[string]string{
	"entry_pagination": `{{ define "entry_pagination" }}
<div class="pagination">
    <div class="pagination-prev">
        {{ if .prevEntry }}
            <a href="{{ .prevEntryRoute }}{{ if .searchQuery }}?q={{ .searchQuery }}{{ end }}" title="{{ .prevEntry.Title }}" data-page="previous" rel="prev">{{ t "pagination.previous" }}</a>
        {{ else }}
            {{ t "pagination.previous" }}
        {{ end }}
    </div>

    <div class="pagination-next">
        {{ if .nextEntry }}
            <a href="{{ .nextEntryRoute }}{{ if .searchQuery }}?q={{ .searchQuery }}{{ end }}" title="{{ .nextEntry.Title }}" data-page="next" rel="next">{{ t "pagination.next" }}</a>
        {{ else }}
            {{ t "pagination.next" }}
        {{ end }}
    </div>
</div>
{{ end }}
`,
	"feed_list": `{{ define "feed_list" }}
    <div class="items">
        {{ range .feeds }}
        <article class="item {{ if ne .ParsingErrorCount 0 }}feed-parsing-error{{ end }}">
            <div class="item-header" dir="auto">
                <span class="item-title">
                    {{ if .Icon }}
                        <img src="{{ route "icon" "iconID" .Icon.IconID }}" width="16" height="16" loading="lazy" alt="{{ .Title }}">
                    {{ end }}
                    {{ if .Disabled }} 🚫 {{ end }}
                    <a href="{{ route "feedEntries" "feedID" .ID }}">{{ .Title }}</a>
                </span>
                <span class="feed-entries-counter">
                    (<span title="{{ t "page.feeds.unread_counter" }}">{{ .UnreadCount }}</span>/<span title="{{ t "page.feeds.read_counter" }}">{{ .ReadCount }}</span>)
                </span>
                <span class="category">
                    <a href="{{ route "categoryEntries" "categoryID" .Category.ID }}">{{ .Category.Title }}</a>
                </span>
            </div>
            <div class="item-meta">
                <ul class="item-meta-info">
                    <li dir="auto">
                        <a href="{{ .SiteURL | safeURL  }}" title="{{ .SiteURL }}" target="_blank" rel="noopener noreferrer" referrerpolicy="no-referrer" data-original-link="true">{{ domain .SiteURL }}</a>
                    </li>
                    <li>
                        {{ t "page.feeds.last_check" }} <time datetime="{{ isodate .CheckedAt }}" title="{{ isodate .CheckedAt }}">{{ elapsed $.user.Timezone .CheckedAt }}</time>
                    </li>
                </ul>
                <ul class="item-meta-icons">
                    <li>
                        <a href="{{ route "refreshFeed" "feedID" .ID }}">{{ template "icon_refresh" }}<span class="icon-label">{{ t "menu.refresh_feed" }}</span></a>
                    </li>
                    <li>
                        <a href="{{ route "editFeed" "feedID" .ID }}">{{ template "icon_edit" }}<span class="icon-label">{{ t "menu.edit_feed" }}</span></a>
                    </li>
                    <li>
                        <a href="#"
                            data-confirm="true"
                            data-label-question="{{ t "confirm.question" }}"
                            data-label-yes="{{ t "confirm.yes" }}"
                            data-label-no="{{ t "confirm.no" }}"
                            data-label-loading="{{ t "confirm.loading" }}"
                            data-url="{{ route "removeFeed" "feedID" .ID }}">{{ template "icon_delete" }}<span class="icon-label">{{ t "action.remove" }}</span></a>
                    </li>
                    {{ if .UnreadCount }}
                      <li>
                        <a href="{{ route "markFeedAsRead" "feedID" .ID }}">{{ template "icon_read" }}<span class="icon-label">{{ t "menu.mark_all_as_read" }}</span></a>
                      </li>
                    {{ end }}
                </ul>
            </div>
            {{ if ne .ParsingErrorCount 0 }}
                <div class="parsing-error">
                    <strong title="{{ .ParsingErrorMsg }}" class="parsing-error-count">{{ plural "page.feeds.error_count" .ParsingErrorCount .ParsingErrorCount }}</strong>
                    - <small class="parsing-error-message">{{ .ParsingErrorMsg }}</small>
                </div>
            {{ end }}
        </article>
        {{ end }}
    </div>
{{ end }}
`,
	"feed_menu": `{{ define "feed_menu" }}
<ul>
    <li>
        <a href="{{ route "feeds" }}">{{ t "menu.feeds" }}</a>
    </li>
    <li>
        <a href="{{ route "addSubscription" }}">{{ t "menu.add_feed" }}</a>
    </li>
    <li>
        <a href="{{ route "export" }}">{{ t "menu.export" }}</a>
    </li>
    <li>
        <a href="{{ route "import" }}">{{ t "menu.import" }}</a>
    </li>
    <li>
        <a href="{{ route "refreshAllFeeds" }}">{{ t "menu.refresh_all_feeds" }}</a>
    </li>
</ul>
{{ end }}`,
	"icons": `<!--

MIT License

Copyright (c) 2020 Paweł Kuna

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

-->
{{ define "icon_read" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-circle-check" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <circle cx="12" cy="12" r="9" />
    <path d="M9 12l2 2l4 -4" />
</svg>
{{ end }}
{{ define "icon_unread" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-circle-x" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <circle cx="12" cy="12" r="9" />
    <path d="M10 10l4 4m0 -4l-4 4" />
</svg>
{{ end }}
{{ define "icon_star" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-star" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M12 17.75l-6.172 3.245 1.179-6.873-4.993-4.867 6.9-1.002L12 2l3.086 6.253 6.9 1.002-4.993 4.867 1.179 6.873z" />
</svg>
{{ end }}
{{ define "icon_unstar" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-unstar" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path fill="currentColor" d="M12 17.75l-6.172 3.245 1.179-6.873-4.993-4.867 6.9-1.002L12 2l3.086 6.253 6.9 1.002-4.993 4.867 1.179 6.873z" />
</svg>
{{ end }}
{{ define "icon_save" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-download" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M4 17v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2 -2v-2" />
    <polyline points="7 11 12 16 17 11" />
    <line x1="12" y1="4" x2="12" y2="16" />
</svg>
{{ end }}
{{ define "icon_scraper" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-cloud-download" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M19 18a3.5 3.5 0 0 0 0 -7h-1a5 4.5 0 0 0 -11 -2a4.6 4.4 0 0 0 -2.1 8.4" />
    <line x1="12" y1="13" x2="12" y2="22" />
    <polyline points="9 19 12 22 15 19" />
</svg>
{{ end }}
{{ define "icon_share" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-share" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <circle cx="6" cy="12" r="3" />
    <circle cx="18" cy="6" r="3" />
    <circle cx="18" cy="18" r="3" />
    <line x1="8.7" y1="10.7" x2="15.3" y2="7.3" />
    <line x1="8.7" y1="13.3" x2="15.3" y2="16.7" />
</svg>
{{ end }}
{{ define "icon_comment" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-message-circle" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M3 20l1.3 -3.9a9 8 0 1 1 3.4 2.9l-4.7 1" />
    <line x1="12" y1="12" x2="12" y2="12.01" />
    <line x1="8" y1="12" x2="8" y2="12.01" />
    <line x1="16" y1="12" x2="16" y2="12.01" />
</svg>
{{ end }}
{{ define "icon_external_link" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-external-link" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M11 7h-5a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-5" />
    <line x1="10" y1="14" x2="20" y2="4" />
    <polyline points="15 4 20 4 20 9" />
</svg>
{{ end }}
{{ define "icon_delete" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-trash" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <line x1="4" y1="7" x2="20" y2="7" />
    <line x1="10" y1="11" x2="10" y2="17" />
    <line x1="14" y1="11" x2="14" y2="17" />
    <path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
    <path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
</svg>
{{ end }}
{{ define "icon_edit" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-edit" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M9 7 h-3a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-3" />
    <path d="M9 15h3l8.5 -8.5a1.5 1.5 0 0 0 -3 -3l-8.5 8.5v3" />
    <line x1="16" y1="5" x2="19" y2="8" />
</svg>
{{ end }}
{{ define "icon_feeds" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-folders" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M9 4h3l2 2h5a2 2 0 0 1 2 2v7a2 2 0 0 1 -2 2h-10a2 2 0 0 1 -2 -2v-9a2 2 0 0 1 2 -2" />
    <path d="M17 17v2a2 2 0 0 1 -2 2h-10a2 2 0 0 1 -2 -2v-9a2 2 0 0 1 2 -2h2" />
</svg>
{{ end }}
{{ define "icon_entries" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-news" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M16 6h3a1 1 0 0 1 1 1v11a2 2 0 0 1 -4 0v-13a1 1 0 0 0 -1 -1h-10a1 1 0 0 0 -1 1v12a3 3 0 0 0 3 3h11" />
    <line x1="8" y1="8" x2="12" y2="8" />
    <line x1="8" y1="12" x2="12" y2="12" />
    <line x1="8" y1="16" x2="12" y2="16" />
</svg>
{{ end }}
{{ define "icon_refresh" }}
<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-refresh" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
    <path stroke="none" d="M0 0h24v24H0z"/>
    <path d="M20 11a8.1 8.1 0 0 0 -15.5 -2m-.5 -5v5h5" />
    <path d="M4 13a8.1 8.1 0 0 0 15.5 2m.5 5v-5h-5" />
</svg>
{{ end }}`,
	"item_meta": `{{ define "item_meta" }}
<div class="item-meta">
    <ul class="item-meta-info">
        <li>
            <a href="{{ route "feedEntries" "feedID" .entry.Feed.ID }}" title="{{ .entry.Feed.SiteURL }}" data-feed-link="true">{{ truncate .entry.Feed.Title 35 }}</a>
        </li>
        <li>
            <time datetime="{{ isodate .entry.Date }}" title="{{ isodate .entry.Date }}">{{ elapsed .user.Timezone .entry.Date }}</time>
        </li>
        {{ if and .user.ShowReadingTime (gt .entry.ReadingTime 0) }}
        <li>
            <span>
            {{ plural "entry.estimated_reading_time" .entry.ReadingTime .entry.ReadingTime }}
            </span>
        </li>
        {{ end }}
    </ul>
    <ul class="item-meta-icons">
        <li>
            <a href="#"
                title="{{ t "entry.status.title" }}"
                data-toggle-status="true"
                data-label-read="{{ t "entry.status.read" }}"
                data-label-unread="{{ t "entry.status.unread" }}"
                data-value="{{ if eq .entry.Status "read" }}read{{ else }}unread{{ end }}"
                >{{ if eq .entry.Status "read" }}{{ template "icon_unread" }}{{ else }}{{ template "icon_read" }}{{ end }}<span class="icon-label">{{ if eq .entry.Status "read" }}{{ t "entry.status.unread" }}{{ else }}{{ t "entry.status.read" }}{{ end }}</span></a>
        </li>
        <li>
            <a href="#"
                data-toggle-bookmark="true"
                data-bookmark-url="{{ route "toggleBookmark" "entryID" .entry.ID }}"
                data-label-loading="{{ t "entry.state.saving" }}"
                data-label-star="{{ t "entry.bookmark.toggle.on" }}"
                data-label-unstar="{{ t "entry.bookmark.toggle.off" }}"
                data-value="{{ if .entry.Starred }}star{{ else }}unstar{{ end }}"
                >{{ if .entry.Starred }}{{ template "icon_unstar" }}{{ else }}{{ template "icon_star" }}{{ end }}<span class="icon-label">{{ if .entry.Starred }}{{ t "entry.bookmark.toggle.off" }}{{ else }}{{ t "entry.bookmark.toggle.on" }}{{ end }}</span></a>
        </li>
        {{ if .entry.ShareCode }}
            <li>
                <a href="{{ route "sharedEntry" "shareCode" .entry.ShareCode }}"
                    title="{{ t "entry.shared_entry.title" }}"
                    target="_blank">{{ template "icon_share" }}<span class="icon-label">{{ t "entry.shared_entry.label" }}</span></a>
            </li>
        {{ end }}
        {{ if .hasSaveEntry }}
            <li>
                <a href="#"
                    title="{{ t "entry.save.title" }}"
                    data-save-entry="true"
                    data-save-url="{{ route "saveEntry" "entryID" .entry.ID }}"
                    data-label-loading="{{ t "entry.state.saving" }}"
                    data-label-done="{{ t "entry.save.completed" }}"
                    >{{ template "icon_save" }}<span class="icon-label">{{ t "entry.save.label" }}</span></a>
            </li>
        {{ end }}
        <li>
            <a href="{{ .entry.URL | safeURL  }}"
                target="_blank"
                rel="noopener noreferrer"
                referrerpolicy="no-referrer"
                data-original-link="true">{{ template "icon_external_link" }}<span class="icon-label">{{ t "entry.external_link.label" }}</span></a>
        </li>
        {{ if .entry.CommentsURL }}
            <li>
                <a href="{{ .entry.CommentsURL | safeURL  }}"
                    title="{{ t "entry.comments.title" }}"
                    target="_blank"
                    rel="noopener noreferrer"
                    referrerpolicy="no-referrer"
                    data-comments-link="true">{{ template "icon_comment" }}<span class="icon-label">{{ t "entry.comments.label" }}</span></a>
            </li>
        {{ end }}
    </ul>
</div>
{{ end }}
`,
	"layout": `{{ define "base" }}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>{{template "title" .}} - Miniflux</title>

    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=no">
    <meta name="mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-title" content="Miniflux">
    <link rel="manifest" href="{{ route "webManifest" }}" crossorigin="use-credentials"/>

    <meta name="robots" content="noindex,nofollow">
    <meta name="referrer" content="no-referrer">
    <meta name="google" content="notranslate">

    <!-- Favicons -->
    <link rel="icon" type="image/png" sizes="16x16" href="{{ route "appIcon" "filename" "favicon-16.png" }}">
    <link rel="icon" type="image/png" sizes="32x32" href="{{ route "appIcon" "filename" "favicon-32.png" }}">

    <!-- Android icons -->
    <link rel="icon" type="image/png" sizes="128x128" href="{{ route "appIcon" "filename" "icon-128.png" }}">
    <link rel="icon" type="image/png" sizes="192x192" href="{{ route "appIcon" "filename" "icon-192.png" }}">

    <!-- iOS icons -->
    <link rel="apple-touch-icon" sizes="120x120" href="{{ route "appIcon" "filename" "icon-120.png" }}">
    <link rel="apple-touch-icon" sizes="152x152" href="{{ route "appIcon" "filename" "icon-152.png" }}">
    <link rel="apple-touch-icon" sizes="167x167" href="{{ route "appIcon" "filename" "icon-167.png" }}">
    <link rel="apple-touch-icon" sizes="180x180" href="{{ route "appIcon" "filename" "icon-180.png" }}">

    {{ if .csrf }}
        <meta name="X-CSRF-Token" value="{{ .csrf }}">
    {{ end }}

    <meta name="theme-color" content="{{ theme_color .theme }}">
    <link rel="stylesheet" type="text/css" href="{{ route "stylesheet" "name" .theme }}?{{ .theme_checksum }}">
    {{ if and .user .user.Stylesheet }}
    <link rel="stylesheet" type="text/css" href="{{ route "stylesheet" "name" "custom_css" }}">
    {{ end }}

    <script type="text/javascript" src="{{ route "javascript" "name" "app" }}?{{ .app_js_checksum }}" defer></script>
    <script type="text/javascript" src="{{ route "javascript" "name" "service-worker" }}?{{ .sw_js_checksum }}" defer id="service-worker-script"></script>
</head>
<body
    data-entries-status-url="{{ route "updateEntriesStatus" }}"
    data-refresh-all-feeds-url="{{ route "refreshAllFeeds" }}"
    {{ if .user }}{{ if not .user.KeyboardShortcuts }}data-disable-keyboard-shortcuts="true"{{ end }}{{ end }}>
    <div class="toast-wrap">
        <span class="toast-msg"></span>
    </div>
    {{ if .user }}
    <header class="header">
        <nav>
            <div class="logo">
                <a href="{{ route "unread" }}">Mini<span>flux</span></a>
            </div>
            <ul>
                <li {{ if eq .menu "unread" }}class="active"{{ end }} title="{{ t "tooltip.keyboard_shortcuts" "g u" }}">
                    <a href="{{ route "unread" }}" data-page="unread">{{ t "menu.unread" }}
                      {{ if gt .countUnread 0 }}
                          <span class="unread-counter-wrapper">(<span class="unread-counter">{{ .countUnread }}</span>)</span>
                      {{ end }}
                    </a>
                </li>
                <li {{ if eq .menu "starred" }}class="active"{{ end }} title="{{ t "tooltip.keyboard_shortcuts" "g b" }}">
                    <a href="{{ route "starred" }}" data-page="starred">{{ t "menu.starred" }}</a>
                </li>
                <li {{ if eq .menu "history" }}class="active"{{ end }} title="{{ t "tooltip.keyboard_shortcuts" "g h" }}">
                    <a href="{{ route "history" }}" data-page="history">{{ t "menu.history" }}</a>
                </li>
                <li {{ if eq .menu "feeds" }}class="active"{{ end }} title="{{ t "tooltip.keyboard_shortcuts" "g f" }}">
                    <a href="{{ route "feeds" }}" data-page="feeds">{{ t "menu.feeds" }}
                      {{ if gt .countErrorFeeds 0 }}
                          <span class="error-feeds-counter-wrapper">(<span class="error-feeds-counter">{{ .countErrorFeeds }}</span>)</span>
                      {{ end }}
                    </a>
                </li>
                <li {{ if eq .menu "categories" }}class="active"{{ end }} title="{{ t "tooltip.keyboard_shortcuts" "g c" }}">
                    <a href="{{ route "categories" }}" data-page="categories">{{ t "menu.categories" }}</a>
                </li>
                <li {{ if eq .menu "settings" }}class="active"{{ end }} title="{{ t "tooltip.keyboard_shortcuts" "g s" }}">
                    <a href="{{ route "settings" }}" data-page="settings">{{ t "menu.settings" }}</a>
                </li>
                <li>
                    <a href="{{ route "logout" }}" title="{{ t "tooltip.logged_user" .user.Username }}">{{ t "menu.logout" }}</a>
                </li>
            </ul>
            <div class="search">
                <div class="search-toggle-switch {{ if $.searchQuery }}has-search-query{{ end }}">
                    <a href="#" data-action="search">&laquo;&nbsp;{{ t "search.label" }}</a>
                </div>
                <form action="{{ route "searchEntries" }}" class="search-form {{ if $.searchQuery }}has-search-query{{ end }}">
                    <input type="search" name="q" id="search-input" placeholder="{{ t "search.placeholder" }}" {{ if $.searchQuery }}value="{{ .searchQuery }}"{{ end }} required>
                </form>
            </div>
        </nav>
    </header>
    {{ end }}
    {{ if .flashMessage }}
        <div class="flash-message alert alert-success">{{ .flashMessage }}</div>
    {{ end }}
    {{ if .flashErrorMessage }}
        <div class="flash-error-message alert alert-error">{{ .flashErrorMessage }}</div>
    {{ end }}
    <main>
        {{template "content" .}}
    </main>
    <template id="keyboard-shortcuts">
        <div id="modal-left">
            <a href="#" class="btn-close-modal">x</a>
            <h3>{{ t "page.keyboard_shortcuts.title" }}</h3>

            <div class="keyboard-shortcuts">
                <p>{{ t "page.keyboard_shortcuts.subtitle.sections" }}</p>
                <ul>
                    <li>{{ t "page.keyboard_shortcuts.go_to_unread" }} = <strong>g + u</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_starred" }} = <strong>g + b</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_history" }} = <strong>g + h</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_feeds" }} = <strong>g + f</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_categories" }} = <strong>g + c</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_settings" }} = <strong>g + s</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.show_keyboard_shortcuts" }} = <strong>?</strong></li>
                </ul>

                <p>{{ t "page.keyboard_shortcuts.subtitle.items" }}</p>
                <ul>
                    <li>{{ t "page.keyboard_shortcuts.go_to_previous_item" }} = <strong>p</strong>, <strong>k</strong>, <strong>◄</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_next_item" }} = <strong>n</strong>, <strong>j</strong>, <strong>►</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_feed" }} = <strong>F</strong></li>
                </ul>

                <p>{{ t "page.keyboard_shortcuts.subtitle.pages" }}</p>
                <ul>
                    <li>{{ t "page.keyboard_shortcuts.go_to_previous_page" }} = <strong>h</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_next_page" }} = <strong>l</strong></li>
                </ul>

                <p>{{ t "page.keyboard_shortcuts.subtitle.actions" }}</p>
                <ul>
                    <li>{{ t "page.keyboard_shortcuts.open_item" }} = <strong>o</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.open_original" }} = <strong>v</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.open_original_same_window" }} = <strong>V</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.open_comments" }} = <strong>c</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.open_comments_same_window" }} = <strong>C</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.toggle_read_status" }} = <strong>m</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.mark_page_as_read" }} = <strong>A</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.download_content" }} = <strong>d</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.toggle_bookmark_status" }} = <strong>f</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.save_article" }} = <strong>s</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.scroll_item_to_top" }} = <strong>z + t</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.refresh_all_feeds" }} = <strong>R</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.remove_feed" }} = <strong>#</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.go_to_search" }} = <strong>/</strong></li>
                    <li>{{ t "page.keyboard_shortcuts.close_modal" }} = <strong>Esc</strong></li>
                </ul>
            </div>
        </div>
    </template>
    <template id="icon_read">
        {{ template "icon_read" }}
    </template>
    <template id="icon_unread">
        {{ template "icon_unread" }}
    </template>
    <template id="icon_star">
        {{ template "icon_star" }}
    </template>
    <template id="icon_unstar">
        {{ template "icon_unstar" }}
    </template>
</body>
</html>
{{ end }}
`,
	"pagination": `{{ define "pagination" }}
<div class="pagination">
    <div class="pagination-prev">
        {{ if .ShowPrev }}
            <a href="{{ .Route }}{{ if gt .PrevOffset 0 }}?offset={{ .PrevOffset }}{{ if .SearchQuery }}&amp;q={{ .SearchQuery }}{{ end }}{{ else }}{{ if .SearchQuery }}?q={{ .SearchQuery }}{{ end }}{{ end }}" data-page="previous" rel="prev">{{ t "pagination.previous" }}</a>
        {{ else }}
            {{ t "pagination.previous" }}
        {{ end }}
    </div>

    <div class="pagination-next">
        {{ if .ShowNext }}
            <a href="{{ .Route }}?offset={{ .NextOffset }}{{ if .SearchQuery }}&amp;q={{ .SearchQuery }}{{ end }}" data-page="next" rel="next">{{ t "pagination.next" }}</a>
        {{ else }}
            {{ t "pagination.next" }}
        {{ end }}
    </div>
</div>
{{ end }}
`,
	"settings_menu": `{{ define "settings_menu" }}
<ul>
    <li>
        <a href="{{ route "settings" }}">{{ t "menu.settings" }}</a>
    </li>
    <li>
        <a href="{{ route "integrations" }}">{{ t "menu.integrations" }}</a>
    </li>
    <li>
        <a href="{{ route "apiKeys" }}">{{ t "menu.api_keys" }}</a>
    </li>
    <li>
        <a href="{{ route "sessions" }}">{{ t "menu.sessions" }}</a>
    </li>
    {{ if .user.IsAdmin }}
        <li>
            <a href="{{ route "users" }}">{{ t "menu.users" }}</a>
        </li>
    {{ end }}
    <li>
        <a href="{{ route "about" }}">{{ t "menu.about" }}</a>
    </li>
</ul>
{{ end }}`,
}

var templateCommonMapChecksums = map[string]string{
	"entry_pagination": "cdca9cf12586e41e5355190b06d9168f57f77b85924d1e63b13524bc15abcbf6",
	"feed_list":        "931e43d328a116318c510de5658c688cd940b934c86b6ec82a472e1f81e020ae",
	"feed_menu":        "318d8662dda5ca9dfc75b909c8461e79c86fb5082df1428f67aaf856f19f4b50",
	"icons":            "7161afa4cce46245a99cb1e49a605d3ff30e907c3f568ef9c17218718d20e042",
	"item_meta":        "fefa219c8296f0370632336ed59a2c8b0c2146ee77f3b10de1d9b87982219dc5",
	"layout":           "03c77ed0163b790c0622ecec173119537087c66f6a3925a931ae83a9a94d32cf",
	"pagination":       "7b61288e86283c4cf0dc83bcbf8bf1c00c7cb29e60201c8c0b633b2450d2911f",
	"settings_menu":    "e2b777630c0efdbc529800303c01d6744ed3af80ec505ac5a5b3f99c9b989156",
}