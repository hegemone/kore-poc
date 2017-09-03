module Kore
  module Comm
    module Plugin
      module ClassMethods
        attr_reader :matchers
        def match(pattern, opts)
          @matchers[pattern] = opts[:method]
        end
        def self.extended(base)
          base.instance_exec do
            @matchers = {}
          end
        end
      end
      module InstanceMethods
        attr_accessor :engine
        def handle_ingress(msg)
          EM.defer do
            self.class.matchers.each do |pattern, f|
              match = pattern.match(msg.content)
              if !!match
                method(f).call(msg, match)
              end
            end
          end
        end
        def name
          raise "ERROR: Kore::Comm::Plugin must implement 'name'"
        end
        def send_message(originator, msg)
          @engine.handle_egress(originator, msg)
        end
      end
      def self.included(base)
        base.send :include, InstanceMethods
        base.extend ClassMethods
      end
    end
  end
end
