package main

const indexHTMLTemplate = `
<!DOCTYPE html>
<html>
   <head>
      <title>Node.js App</title>
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <link rel="shortcut icon" href="/shortcut-icon.png">
      <link rel="stylesheet" href="/pui-3.0.0/pivotal-ui.min.css">
      <link rel="stylesheet" href="/font-mfizz-2.3.0/font-mfizz.css">
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.2.2/jquery.min.js"></script><script src="/javascripts/ping.js"></script>
   </head>
   <body>
      <div class="container pan">
         <section class="header">
            <div class="row man">
               <div class="col-xs-24 col-sm-24 col-md-24 col-lg-24 mvxxl txt-c"><a href="https://cloudfoundry.org" target="_blank"><img class="txt-b" src="/images/the-rabbit.svg" width="52" alt="Cloud Foundry" title="Cloud Foundry"></a><span class="h1 em-high txt-m mhl type-neutral-4"> + </span><a href="https://nodejs.org" target="_blank"><img class="txt-m" src="/images/nodejs_logo.png" width="46" alt="Node.js" title="Node.js"></a></div>
            </div>
         </section>
         <section class="content">
            <ul class="list-group">
               <!-- buildpack-->
               <li class="list-group-item pln">
                  <div class="row man">
                     <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3 phxl"><span class="h1 pan man icon-nodejs type-neutral-5" style="top: 4px; left: 2px; position: relative;"></span></div>
                     <div class="col-xs-12 col-sm-15 col-md-15 col-lg-15">
                        <div class="type-neutral-4 small">
                           <div>BUILDPACK</div>
                        </div>
                        <div class="type-ellipsis">Go</div>
                     </div>
                     <div class="col-xs-9 col-sm-6 col-md-6 col-lg-6 ptl txt-r type-ellipsis"><a class="type-accent-4" href="https://docs.cloudfoundry.org/buildpacks/" target="_blank"><span class="small">Buildpacks</span><span class="fa fa-icon fa-external-link plm txt-m small"></span></a></div>
                  </div>
               </li>
               <!-- app-->
               <li class="list-group-item pln">
                  <div class="row man">
                     <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3 phxl"><span class="h1 pan man fa fa-icon fa-play-circle type-neutral-5" style="left: 0px; position: relative;"></span></div>
                     <div class="col-xs-5 col-sm-5 col-md-5 col-lg-5">
                        <div class="type-neutral-4 small">
                           <div class="type-ellipsis">APP NAME</div>
                        </div>
                        <div class="type-ellipsis">{{.AppName}}</div>
                     </div>
                     <div class="col-xs-5 col-sm-5 col-md-10 col-lg-10">
                        <div class="type-neutral-4 small">
                           <div class="type-ellipsis">APP URIS</div>
                        </div>
                        <div class="type-ellipsis">{{.AppUris}}</div>
                     </div>
                     <div class="col-xs-11 col-sm-11 col-md-6 col-lg-6 ptl txt-r type-ellipsis"><a class="type-accent-4" href="https://docs.cloudfoundry.org/devguide/deploy-apps/routes-domains.html" target="_blank"><span class="small">Routes & Domains</span><span class="fa fa-icon fa-external-link plm txt-m small"></span></a></div>
                  </div>
               </li>
               <!-- limits-->
               <li class="list-group-item pln">
                  <div class="row man">
                     <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3 phxl"><span class="h1 pan man fa fa-icon fa-dashboard type-neutral-5" style="left: -2px; position: relative;"></span></div>
                     <div class="col-xs-5 col-sm-5 col-md-5 col-lg-5">
                        <div class="type-neutral-4 small">
                           <div class="type-ellipsis">INSTANCE INDEX</div>
                        </div>
                        <div class="type-ellipsis">{{.AppInstanceIndex}}</div>
                     </div>
                     <div class="col-xs-5 col-sm-5 col-md-5 col-lg-5">
                        <div class="type-neutral-4 small">
                           <div class="type-ellipsis">MEMORY LIMIT</div>
                        </div>
                        <div class="type-ellipsis">{{.Limits.Mem}}</div>
                     </div>
                     <div class="col-xs-5 col-sm-5 col-md-5 col-lg-5">
                        <div class="type-neutral-4 small">
                           <div class="type-ellipsis">DISK LIMIT</div>
                        </div>
                        <div class="type-ellipsis">{{.Limits.Disk}}</div>
                     </div>
                     <div class="col-xs-6 col-sm-6 col-md-6 col-lg-6 ptl txt-r type-ellipsis"><a class="type-accent-4" href="https://docs.cloudfoundry.org/devguide/deploy-apps/cf-scale.html" target="_blank"><span class="small">Scaling</span><span class="fa fa-icon fa-external-link plm txt-m small"></span></a></div>
                  </div>
               </li>
               <!-- space-->
               <li class="list-group-item pln">
                  <div class="row man">
                     <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3 phxl"><span class="h1 pan man fa fa-icon fa-bullseye type-neutral-5" style="left: 0px; position: relative;"></span></div>
                     <div class="col-xs-12 col-sm-15 col-md-15 col-lg-15">
                        <div class="type-neutral-4 small"><span class="type-ellipsis">SPACE NAME</span></div>
                        <div class="type-ellipsis">{{.AppSpaceName}}</div>
                     </div>
                     <div class="col-xs-9 col-sm-6 col-md-6 col-lg-6 ptl txt-r type-ellipsis"><a class="type-accent-4" href="https://docs.cloudfoundry.org/concepts/roles.html" target="_blank"><span class="small txt-m">Orgs & Spaces</span><span class="fa fa-icon fa-external-link plm txt-m small"></span></a></div>
                  </div>
               </li>
               <!-- services-->
               <li class="list-group-item pln">
                  <div class="row man">
                     <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3 phxl"><span class="h1 pan man fa fa-icon fa-database type-neutral-5" style="left: 0px; position: relative;"></span></div>
                     <!-- no service bound to the app-->
                     <div class="col-xs-12 col-sm-15 col-md-15 col-lg-15">
                        <div class="ptl type-ellipsis">There aren't any services bound to this app.</div>
                     </div>
                     <div class="col-xs-9 col-sm-6 col-md-6 col-lg-6 ptl txt-r type-ellipsis"><a class="type-accent-4" href="https://docs.cloudfoundry.org/devguide/services/managing-services.html" target="_blank"><span class="small">Manage Services</span><span class="fa fa-icon fa-external-link plm txt-m small"></span></a></div>
                  </div>
               </li>
            </ul>
         </section>
         <section class="footer">
            <div class="row mhn mvxxl txt-c">
               <a href="https://docs.cloudfoundry.org/devguide/" target="_blank"><img src="/images/cloud-foundry.svg" width="28" alt="Cloud Foundry" title="Cloud Foundry"><br><span class="txt-c small type-neutral-4">This is a Cloud Foundry sample application.</span></a>
               <div data-ping="ping" style="display: none;"></div>
            </div>
         </section>
      </div>
   </body>
</html>
`
